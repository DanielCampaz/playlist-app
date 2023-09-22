import {
  GetHeaderWhitToken,
  URLLISTADD,
  URLLISTDELETEBYID,
  URLLISTDELETEMANY,
  URLLISTGETALL,
  URLLISTGETALLCODES,
  URLLISTGETBYID,
  URLLISTNEXT,
  URLLISTPOSTCREATE,
  URLLISTRESETBYID,
  URLLISTUPDATEBYID,
} from "../../const";
import { AddType, List } from "../../types";

export default class ListConnection {
  private constructor() {}

  static async deleteManyTable(ids: (string | number)[]) {
    const response = await fetch(URLLISTDELETEMANY, {
      method: "DELETE",
      headers: GetHeaderWhitToken(true),
      body: JSON.stringify({ ids }),
    });
    const datas = await response.json();
    return datas;
  }

  static async getAll(limit = 10, offset = 0) {
    const response = await fetch(URLLISTGETALL(limit, offset), {
      method: "GET",
      headers: GetHeaderWhitToken(true),
    });
    const datas = await response.json();
    return datas;
  }

  static async getAllCodes(idlist: string | number, limit = 10, offset = 0) {
    const response = await fetch(URLLISTGETALLCODES(idlist, limit, offset), {
      method: "GET",
      headers: GetHeaderWhitToken(true),
    });
    const datas = await response.json();
    return datas;
  }

  static async getById(id: string | number) {
    const response = await fetch(URLLISTGETBYID(id), {
      method: "GET",
      headers: GetHeaderWhitToken(false),
    });
    const datas = await response.json();
    return datas;
  }

  static async postCreate(list: Partial<List>) {
    const response = await fetch(URLLISTPOSTCREATE, {
      method: "POST",
      headers: GetHeaderWhitToken(true),
      body: JSON.stringify(list),
    });
    const datas = await response.json();
    return datas;
  }

  static async putUpdate(id: string | number, list: List) {
    const response = await fetch(URLLISTUPDATEBYID(id), {
      method: "PUT",
      headers: GetHeaderWhitToken(true),
      body: JSON.stringify(list),
    });
    const datas = await response.json();
    return datas;
  }

  static async deleteById(id: string | number) {
    const response = await fetch(URLLISTDELETEBYID(id), {
      method: "DELETE",
      headers: GetHeaderWhitToken(false),
    });
    const datas = await response.json();
    return datas;
  }

  static async resetById(id: string | number) {
    const response = await fetch(URLLISTRESETBYID(id), {
      method: "PUT",
      headers: GetHeaderWhitToken(false),
    });
    const datas = await response.json();
    return datas;
  }

  static async postAdd(
    idlist: string | number,
    iduser: string | number,
    code: AddType
  ) {
    const response = await fetch(URLLISTADD(idlist, iduser), {
      method: "POST",
      headers: GetHeaderWhitToken(true),
      body: JSON.stringify(code),
    });
    const datas = await response.json();
    return datas;
  }

  static async getNext(idlist: string | number, iduser: string | number) {
    const response = await fetch(URLLISTNEXT(idlist, iduser), {
      method: "GET",
      headers: GetHeaderWhitToken(false),
    });
    const datas = await response.json();
    return datas;
  }
}
