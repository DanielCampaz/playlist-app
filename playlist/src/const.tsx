import { Typography } from "@mui/material";
import Link from "@mui/material/Link";
import Storage from "./class/storage";

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

export const GetHeaderWhitToken = (CT: boolean) => {
  const headers = new Headers();
  headers.append("Token", Storage.local.getToken());
  if (CT) {
    headers.append("Content-Type", "application/json");
  }
  // Agregar las cabeceras CORS necesarias
  headers.append(
    "Access-Control-Request-Method",
    "GET, POST, PUT, DELETE, OPTIONS"
  );
  headers.append("Access-Control-Request-Headers", "Content-Type, Token"); // Ajusta esto según tus necesidades

  return headers;
};

export const urlPlay = (id: string) => {
  return `/play/${id}`;
};
export const urlAdd = (id: string) => {
  return `/add/${id}`;
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
export const URLDELETEUSER = (id: string | number) => URLUSER + "/delete/" + id;

// LIST
export const URLLIST = URLAPI + "/list";
export const URLLISTDELETEMANY = URLLIST + "/deletemany";
export const URLLISTGETALL = (limit = 10, offset = 0) => {
  return URLLIST + "/g/all?limit=" + limit + "&offset=" + offset;
};
export const URLLISTGETALLCODES = (
  idlist: string | number,
  limit = 10,
  offset = 0
) => {
  return (
    URLLIST + "/codes/all/" + idlist + "?limit=" + limit + "&offset=" + offset
  );
  //return URLLIST + "/g/all?limit=" + limit + "&offset=" + offset;
};
export const URLLISTGETBYID = (id: string | number) => {
  return URLLIST + "/" + id;
};
export const URLLISTPOSTCREATE = URLLIST + "/create";
export const URLLISTUPDATEBYID = (id: string | number) => {
  return URLLIST + "/update/" + id;
};
export const URLLISTDELETEBYID = (id: string | number) => {
  return URLLIST + "/update/" + id;
};
export const URLLISTRESETBYID = (id: string | number) => {
  return URLLIST + "/restar/" + id;
};
export const URLLISTADD = (
  idlist: string | number,
  iduser: string | number
) => {
  return URLLIST + "/add/" + idlist + "/" + iduser;
};
export const URLLISTNEXT = (
  idlist: string | number,
  iduser: string | number
) => {
  return URLLIST + "/next/" + idlist + "/" + iduser;
};
/*
    Local Storage 
*/
export const SESSIONLOCALST = "sessionPlayList";
export const SESSIONTOKENLOCALST = "sessionTokenPlayList";
