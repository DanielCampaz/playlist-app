export type WID<T> = T & {
  id: string;
};

export interface User {
  name: string;
  lastname: string;
  email: string;
  password: string;
}

export interface List {
  name: string;
  iduser: string;
  act: number;
  counts: number;
}

export interface Code {
  code: string;
  order_number: number;
  isPlatey: boolean;
  idUser: string;
  idList: string;
}

export interface AddType {
  ifr: string;
  type: string;
}

export interface Paginate<T> {
  data: T;
  limit: string;
  offset: string;
}
