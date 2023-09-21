interface Navigation {
  name: string;
  to: string;
  current: boolean;
}

const navigation: Navigation[] = [
  { name: "Home", to: "home", current: true },
  { name: "list's", to: "lists", current: false },
];

export default navigation;
