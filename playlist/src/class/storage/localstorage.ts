import { SESSIONLOCALST, SESSIONTOKENLOCALST } from "../../const";

export default class LocalStorage {
  constructor() {}
  private localstorage = localStorage;
  save(key: string, data: any): boolean {
    this.localstorage.setItem(key, JSON.stringify(data));

    return true;
  }

  get(key: string): null | string {
    const item = this.localstorage.getItem(key);

    if (item !== null) {
      return JSON.parse(item);
    }

    return null;
  }

  saveSession(data: any) {
    if (data.token) {
      this.localstorage.setItem(
        SESSIONTOKENLOCALST,
        JSON.stringify(data.token)
      );
      if (data.data) {
        if (data.data.password) {
          this.localstorage.setItem(
            SESSIONLOCALST,
            JSON.stringify({
              ...data.data,
              password: "",
            })
          );
        } else {
          this.localstorage.setItem(SESSIONLOCALST, JSON.stringify(data.data));
        }
      }
    }
  }

  getToken() {
    const token = this.get(SESSIONTOKENLOCALST);
    if (token === null) {
      return "";
    } else {
      return token;
    }
  }

  getSession() {
    const data = this.get(SESSIONLOCALST);
    if (data === null) {
      return {};
    } else {
      return data;
    }
  }

  singOut() {
    this.delete(SESSIONLOCALST);
    return this.delete(SESSIONTOKENLOCALST);
  }

  delete(key: string): boolean {
    this.localstorage.removeItem(key);

    return true;
  }
}
