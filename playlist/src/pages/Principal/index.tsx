import { useEffect } from "react";
import { Outlet, useNavigate } from "react-router-dom";
import Navbar from "../../components/Navbar";

export default function Principal() {
  const navigate = useNavigate();
  useEffect(() => {
    return () => {
      navigate("/home", {
        replace: true,
      });
    };
  }, []);

  return (
    <div>
      <Navbar />
      <Outlet />
    </div>
  );
}
