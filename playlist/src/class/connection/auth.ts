import { URLLOGIN, URLSINGUP } from "../../const";
import { User } from "../../types";

export interface PropsLogin {
  email: string;
  password: string;
}
export default class AuthConnection {
  private constructor() {}

  static async Login(data: PropsLogin) {
    console.log(URLLOGIN, URLSINGUP);
    const response = await fetch(URLLOGIN, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(data),
    });
    const datas = await response.json();
    return datas;
  }
  static async SingUp(data: User) {
    const response = await fetch(URLSINGUP, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(data),
    });
    const datas = await response.json();
    return datas;
  }
}
