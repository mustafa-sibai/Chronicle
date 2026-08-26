package common

type HttpCodes int

const (
	HttpCodes_Unknown             HttpCodes = -1
	HttpCodes_OK                  HttpCodes = 200
	HttpCodes_Created             HttpCodes = 201
	HttpCodes_BadRequest          HttpCodes = 400
	HttpCodes_Unauthorized        HttpCodes = 401
	HttpCodes_Conflict            HttpCodes = 409
	HttpCodes_InternalServerError HttpCodes = 500
)

type ResponseCodes int

const (
	ResponseCodes_Unknown         ResponseCodes = iota - 1 // -1
	ResponseCodes_Register                                 // 0
	ResponseCodes_Login                                    // 1
	ResponseCodes_Refresh                                  // 2
	ResponseCodes_Logout                                   // 3
	ResponseCodes_CreateSession                            // 4
	ResponseCodes_DestroySession                           // 5
	ResponseCodes_UpdateEmail                              // 6
	ResponseCodes_UpdatePassword                           // 7
	ResponseCodes_UpdateUsername                           // 8
	ResponseCodes_CreateCharacter                          // 9
	ResponseCodes_ListCharacters                           // 10
	ResponseCodes_DeleteCharacter                          // 11
)
