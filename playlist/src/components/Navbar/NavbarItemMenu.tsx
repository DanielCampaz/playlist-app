import { Menu } from "@headlessui/react";
import { Link } from "react-router-dom";

export interface PropsNavbarMenu {
  classNames: (...data: any[]) => string;
  name: string;
  to: string;
  func: () => void;
  type: "Link" | "Button";
}

export default function NavbarItemMenu({
  classNames,
  type,
  func,
  name,
  to,
}: PropsNavbarMenu) {
  return (
    <Menu.Item>
      {({ active }) => (
        <Link
          to={to}
          className={classNames(
            active ? "bg-gray-100" : "",
            "block px-4 py-2 text-sm text-gray-700"
          )}
          onClick={(e) => {
            e.preventDefault();
            if (type === "Button") {
              func();
            }
          }}
        >
          {name}
        </Link>
      )}
    </Menu.Item>
  );
}
