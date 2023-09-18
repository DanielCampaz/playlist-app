"use client";
import AuthConnection, { PropsLogin } from "@/class/connection/auth";
import AuthInput, { PropsAuthInput } from "./AuthInput";
import { useForm } from "react-hook-form";
import { User } from "@/types";
import { URLLOGIN, URLSINGUP } from "@/const";

interface Props {
  type: "login" | "signup";
}

const placeholder = "example@gmail.com";

export default function Auth({ type }: Props) {
  const { register, handleSubmit } = useForm();
  const dataLogin: PropsAuthInput[] = [
    {
      id: "email",
      label: "email",
      name: "email",
      placeholder,
      type: "email",
    },
    {
      id: "password",
      label: "password",
      name: "password",
      placeholder: "23132654",
      type: "password",
    },
  ];

  const dataSingUp: PropsAuthInput[] = [
    {
      id: "name",
      label: "name",
      name: "name",
      placeholder: "Jose Luis",
      type: "text",
    },
    {
      id: "lastname",
      label: "lastname",
      name: "lastname",
      placeholder: "Rodriguez",
      type: "text",
    },
    {
      id: "email",
      label: "email",
      name: "email",
      placeholder,
      type: "email",
    },
    {
      id: "password",
      label: "password",
      name: "password",
      placeholder: "23132654",
      type: "password",
    },
  ];

  const handleSubmitData = async (data: any) => {
    if (type === "login") {
      //const responLogin = await AuthConnection.Login(data as PropsLogin);
      console.log(URLLOGIN, URLSINGUP);
      // const response = await fetch(URLLOGIN, {
      //   method: "POST",
      //   headers: {
      //     "Content-Type": "application/json",
      //   },
      //   body: JSON.stringify(data),
      // });
      // TODO: Send Login to
    } else {
      const responSingUp = await AuthConnection.SingUp(data as User);
      // TODO: Send Sing Up to
    }
    console.log(data);
  };

  return (
    <section className="bg-gray-50 dark:bg-gray-900">
      <div className="flex flex-col items-center justify-center px-6 py-8 mx-auto md:h-screen lg:py-0">
        <a
          href="#"
          className="flex items-center mb-6 text-2xl font-semibold text-gray-900 dark:text-white"
        >
          Playlist App
        </a>
        <div className="w-full bg-white rounded-lg shadow dark:border md:mt-0 sm:max-w-md xl:p-0 dark:bg-gray-800 dark:border-gray-700">
          <div className="p-6 space-y-4 md:space-y-6 sm:p-8">
            <h1 className="text-xl font-bold leading-tight tracking-tight text-gray-900 md:text-2xl dark:text-white">
              {type === "login"
                ? "Sign in to your account"
                : "Create to your account"}
            </h1>
            <form
              className="space-y-4 md:space-y-6"
              onSubmit={handleSubmit(handleSubmitData)}
            >
              {type === "login"
                ? dataLogin.map((inp: PropsAuthInput, index) => {
                    const inpp: PropsAuthInput = {
                      ...inp,
                      register,
                    };
                    return <AuthInput key={index} {...inpp} />;
                  })
                : dataSingUp.map((inp: PropsAuthInput, index) => {
                    const inpp: PropsAuthInput = {
                      ...inp,
                      register,
                    };
                    return <AuthInput key={index} {...inpp} />;
                  })}
              <div className="flex items-center justify-between">
                {type === "login" ? (
                  <>
                    <div className="flex items-start">
                      <div className="flex items-center h-5">
                        <input
                          id="remember"
                          aria-describedby="remember"
                          type="checkbox"
                          className="w-4 h-4 border border-gray-300 rounded bg-gray-50 focus:ring-3 focus:ring-primary-300 dark:bg-gray-700 dark:border-gray-600 dark:focus:ring-primary-600 dark:ring-offset-gray-800"
                        />
                      </div>
                      <div className="ml-3 text-sm">
                        <label
                          htmlFor="remember"
                          className="text-gray-500 dark:text-gray-300"
                        >
                          Remember me
                        </label>
                      </div>
                    </div>
                    <a
                      href="#"
                      className="text-sm font-medium text-primary-600 hover:underline dark:text-primary-500"
                    >
                      Forgot password?
                    </a>
                  </>
                ) : (
                  ""
                )}
              </div>
              <button
                type="submit"
                className="w-full text-white bg-primary-600 hover:bg-primary-700 focus:ring-4 focus:outline-none focus:ring-primary-300 font-medium rounded-lg text-sm px-5 py-2.5 text-center dark:bg-primary-600 dark:hover:bg-primary-700 dark:focus:ring-primary-800"
              >
                {type === "login" ? "Sing In" : "Sing Up"}
              </button>
              <p className="text-sm font-light text-gray-500 dark:text-gray-400">
                {type === "login"
                  ? "Don’t have an account yet?"
                  : "Already you have an account?"}{" "}
                <a
                  href={type === "login" ? "/auth/singup" : "/auth/login"}
                  className="font-medium text-primary-600 hover:underline dark:text-primary-500"
                >
                  {type === "login" ? "Sing Up" : "Login"}
                </a>
              </p>
            </form>
          </div>
        </div>
      </div>
    </section>
  );
}
