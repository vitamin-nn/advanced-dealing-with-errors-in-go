#ifndef GET_USER_HANDLER_H
#define GET_USER_HANDLER_H

#include <stdlib.h>
#include "db.h"
#include "marshalers.h"

typedef enum {
    HTTP_ERR_OK = 0, // 200
    HTTP_BAD_REQUEST = 1, // 400
    HTTP_NOT_FOUND = 2, // 404
    HTTP_UNPROCESSABLE_ENTITY = 3, // 422
    HTTP_INTERNAL_ERROR = 4 // 500    
} http_error_t;

const char* const HTTP_ERR_STRS[] = {
    "200 OK",
    "400 Bad Request",
    "404 Not Found",
    "422 Unprocessable Entity",
    "500 Internal Server Error"
};

const char *http_error_str(http_error_t err)
{
    return HTTP_ERR_STRS[err];
}

http_error_t get_user_handler(char *request_data, char **response_data)
{
    //http_error_t err = HTTP_ERR_OK;
    request_t *req = NULL;
    int err_unm = unmarshal_request(request_data, &req);
    if (err_unm == -1) {
        return HTTP_BAD_REQUEST;
    }

    if (req->user_id <= 0) {
        free(req);

        return HTTP_UNPROCESSABLE_ENTITY;
    }

    user_t *user = NULL;
    db_error_t err_db = db_get_user_by_id(req->user_id, &user);
    if (err_db == DB_ERR_NOT_FOUND) {
        free(req);

        return HTTP_NOT_FOUND;
    } else if (err_db == DB_ERR_INTERNAL) {
        free(req);

        return HTTP_INTERNAL_ERROR;
    }

    int err_resp = marshal_response((response_t){user}, response_data);
    if (err_resp == -1) {
        response_data = NULL;
        free(req);
        free(user->email);
        free(user);

        return HTTP_INTERNAL_ERROR;
    }

    return HTTP_ERR_OK;
}

#endif
