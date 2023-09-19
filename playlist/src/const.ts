export const URLAPI = "http://127.0.0.1:8080/api-v1";

// Auth
export const URLAUTH = URLAPI + "/auth";
export const URLLOGIN = URLAUTH + "/login";
export const URLSINGUP = URLAUTH + "/singup";

// USERS
export const URLUSER = URLAPI + "/users";
export const URLGETUSERBYID = (id: string) => URLUSER + "/" + id;
export const URLGETALLUSERS = URLUSER + "/g/all";
export const URLPOSTUSER = URLUSER + "/create";
export const URLPUTUSER = (id: string) => URLUSER + "/update/" + id;
export const URLDELETEUSER = (id: string) => URLUSER + "/delete/" + id;
