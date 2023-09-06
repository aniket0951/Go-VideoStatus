package helper

const USER_REGISTRATION_SUCCESS = "User registration has been successful"

// REQUIRED_PARAMS --------------- DEVELOPER MESSAGES -----------------//
const REQUIRED_PARAMS = "Please provide a required params."
const FETCHED_FAILED = "Failed to fetch the data from server."
const FETCHED_SUCCESS = "Data fetched successfully."

// DATA_INSERTED -------------- END USERS MESSAGES ----------------//
const DATA_INSERTED = "Operation successful."
const DATA_INSERTED_FAILED = "Operation failed, please try again."
const UPDATE_FAILED = "Update failed, please try again"
const UPDATE_SUCCESS = "Update successful."
const PASSWORD_RESET = "Password reset successful."
const FORGOT_PASSWORD = "Forgot password successful."
const LOGIN_SUCCESS = "Login successful."
const DATA_NOT_FOUND = "Data not found."
const DATA_FOUND = "Data found."

// FAILED_PROCESS -------------- COMMON MESSAGES ----------------//
const FAILED_PROCESS = "failed to process the request."
const PAGINATION_INVALID = "Pagination failed, please try again ..."
const DELETE_FAILED = "failed to delete"
const DELETE_SUCCESS = "successfully deleted"
const INVALID_ID = "invalid id pass"
const INVALID_COUNTRY_NAME = "invalid country name found"
const AUTHENTICATION_FAILED = "unauthorised request"
const PERMISSION_DENIED = "permission denied"
const URL_EXPIRED = "location url expired please try with another url"

// REQUEST_HOST -------------- BACKEND DEV VARIABLES -----------//S
const REQUEST_HOST = ""
const CURRENT_IDX = 0
const PREVIOUS_IDX = 0
const TOTALCOUNT int64 = 0

var TOKEN_ID = ""

var USER_TYPE = ""

// DATA --------------- Response Key Words ------------//
const DATA = "data"
const USER_DATA = "user_data"
const PERMISSION_DATA = "permissions"
const USER_PERMISSION = "user_permission"
const MODULE_DATA = "module_data"
const VIDEO_DATA = "video_data"

// PERMISSION_NOT_FOUND --------- PERMISSIONS ------------- //
const PERMISSION_NOT_FOUND = "permission not found"
const PERMISSION_FOUND = "Permission found."
const PERMISSION_UPDATE_FAILED = "permission update failed please try again"
const HAS_NOT_PERMISSION = "You don't have permission to perform this action"

// ------------------  VIDEO STATUS TAG FOR UPLOADED BY ADMIN --------------//
const VIDEO_INIT = "VIDEO_INIT"                               // WHEN ADMIN UPLOADED A NEW VIDEO
const VIDEO_VERIFY = "VIDEO_VERIFY"                           // ONECE VIDEO VERIFICATION DONE BY ADMIN
const VIDEO_VIRIFICATION_FAILED = "VIDEO_VIRIFICATION_FAILED" // IF VIDEO VERIFICATION FAILED IN VERIFICATION PROCESS
const VIDEO_PUBLISHED = "VIDEO_PUBLISHED"                     // ONECE VIDEO VERIFICATION HAS BEEN SUCCESS AND VIDEO NEED TO PUBLIC
const VIDEO_UNPUBLISHED = "VIDEO_UNPUBLISHED"                 // IF VIDEO HAS BEEN PUBLISH BY MISTAKE THEN IT SHOULD BE UNPUBLISHED

const LOCAL_ADDRESS = "http://localhost:8080/static/" // to fetch video files from storage
