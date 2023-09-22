import { Outlet, createBrowserRouter } from "react-router-dom";
import ErrorPage from "./components/ErrorPage";
import Auth from "./components/Auth/Auth";
import Home from "./pages/Home";
import Principal from "./pages/Principal";
import Add from "./pages/Add";
import Play from "./pages/Play";
import CreateList from "./pages/CreateList";

const Router = createBrowserRouter([
  {
    path: "/",
    element: <Principal />,
    children: [
      {
        path: "home",
        element: <Home />,
      },
      {
        path: "add/:id",
        element: <Add />,
      },
      {
        path: "play/:id",
        element: <Play />,
      },
      {
        path: "create",
        element: <CreateList />,
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
