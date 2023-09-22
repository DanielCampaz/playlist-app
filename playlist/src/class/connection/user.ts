import { GetHeaderWhitToken, URLDELETEUSER } from "../../const";

export default class UserConnection {
  private constructor() {}
  static async deleteById(id: string | number) {
    const response = await fetch(URLDELETEUSER(id), {
      method: "DELETE",
      headers: GetHeaderWhitToken(false),
    });
    const datas = await response.json();
    return datas;
  }
}
