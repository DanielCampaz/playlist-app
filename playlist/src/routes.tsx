import { Outlet, createBrowserRouter } from "react-router-dom";
import ErrorPage from "./components/ErrorPage";
import Auth from "./components/Auth/Auth";
import Home from "./pages/Home";

const Router = createBrowserRouter([
  {
    path: "/",
    element: <Home />,
    errorElement: <ErrorPage />,
  },
  {
    path: "/auth",
    element: <Outlet />,
    children: [
      {
        path: "login",
        element: <Auth type="login" />,
      },
      {
        path: "signup",
        element: <Auth type="signup" />,
      },
    ],
    errorElement: <ErrorPage />,
  },
]);

export default Router;
