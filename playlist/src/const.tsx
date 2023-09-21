import { Typography } from "@mui/material";
import Link from "@mui/material/Link";

export const classNames = (...classes: any[]) => {
  return classes.filter(Boolean).join(" ");
};

export const Copyright = (props: any) => {
  return (
    <Typography
      variant="body2"
      color="text.secondary"
      align="center"
      {...props}
    >
      {"Copyright © "}
      <Link color="inherit" href="/home">
        PlayList App
      </Link>{" "}
      {new Date().getFullYear()}
      {"."}
    </Typography>
  );
};

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

/*
    Local Storage 
*/

export const SESSIONLOCALST = "sessionPlayList";
export const SESSIONTOKENLOCALST = "sessionTokenPlayList";
