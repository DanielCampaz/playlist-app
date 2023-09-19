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

  delete(key: string): boolean {
    this.localstorage.removeItem(key);

    return true;
  }
}
