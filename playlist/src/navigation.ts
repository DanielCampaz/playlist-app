interface Navigation {
  name: string;
  to: string;
  current: boolean;
}

const navigation: Navigation[] = [{ name: "Home", to: "/home", current: true }];

export default navigation;
