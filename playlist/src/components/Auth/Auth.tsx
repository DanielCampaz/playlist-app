import Login from "./Login";
import SignUp from "./SignUp";

interface Props {
  type: "login" | "singup";
}

export default function Auth({ type }: Props) {
  return <>{type === "login" ? <Login /> : <SignUp />}</>;
}
