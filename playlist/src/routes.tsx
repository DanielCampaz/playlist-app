import { Outlet, createBrowserRouter } from "react-router-dom";
import ErrorPage from "./components/ErrorPage";
import Auth from "./components/Auth/Auth";
import Home from "./pages/Home";
import Principal from "./pages/Principal";

const Router = createBrowserRouter([
  {
    path: "/",
    element: <Principal />,
    children: [
      {
        path: "home",
        element: <Home />,
      },
    ],
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
        path: "singup",
        element: <Auth type="singup" />,
      },
    ],
    errorElement: <ErrorPage />,
  },
]);

export default Router;
