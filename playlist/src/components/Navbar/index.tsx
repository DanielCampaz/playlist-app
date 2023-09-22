import { useState } from "react";
import AppBar from "@mui/material/AppBar";
import Box from "@mui/material/Box";
import Toolbar from "@mui/material/Toolbar";
import IconButton from "@mui/material/IconButton";
import Typography from "@mui/material/Typography";
import Menu from "@mui/material/Menu";
import MenuIcon from "@mui/icons-material/Menu";
import Container from "@mui/material/Container";
import Avatar from "@mui/material/Avatar";
import Button from "@mui/material/Button";
import Tooltip from "@mui/material/Tooltip";
import MenuItem from "@mui/material/MenuItem";
import AdbIcon from "@mui/icons-material/Adb";
import pages from "../../navigation";
import Storage from "../../class/storage";
import "./navbar.css";
import UserConnection from "../../class/connection/user";
import { User, WID } from "../../types";
import { ToastContainer, toast } from "react-toastify";

// const settings = ["Profile", "Account", "Dashboard", "Logout"];

export interface NavbarMenu {
  name: string;
  to: string;
  func: () => void;
  type: "Link" | "Button";
}

const settings = () => {
  const wsi: NavbarMenu[] = [
    {
      name: "Login",
      to: "/auth/login",
      func: () => "",
      type: "Link",
    },
    {
      name: "Sing Up",
      to: "/auth/singup",
      func: () => "",
      type: "Link",
    },
  ];

  const wci: NavbarMenu[] = [
    {
      name: "Sing Out",
      to: "",
      func: () => {
        Storage.local.singOut();
      },
      type: "Button",
    },
    {
      name: "Delete Account",
      to: "",
      func: async () => {
        const sesion = Storage.local.getSession() as WID<User> | {};
        if ("id" in sesion) {
          const response = await UserConnection.deleteById(sesion.id);
          if ("error" in response) {
            toast.error(response.message);
          } else if ("message" in response) {
            toast.success(response.message);
          }
        }
      },
      type: "Button",
    },
  ];

  const session = Storage.local.getToken();
  return session === null || session === "" ? wsi : wci;
};

export default function NavBar() {
  const [anchorElNav, setAnchorElNav] = useState<null | HTMLElement>(null);
  const [anchorElUser, setAnchorElUser] = useState<null | HTMLElement>(null);

  const handleOpenNavMenu = (event: React.MouseEvent<HTMLElement>) => {
    setAnchorElNav(event.currentTarget);
  };
  const handleOpenUserMenu = (event: React.MouseEvent<HTMLElement>) => {
    setAnchorElUser(event.currentTarget);
  };

  const handleCloseNavMenu = () => {
    setAnchorElNav(null);
  };

  const handleCloseUserMenu = () => {
    setAnchorElUser(null);
  };

  return (
    <AppBar position="static">
      <Container maxWidth="xl">
        <Toolbar disableGutters>
          <AdbIcon sx={{ display: { xs: "none", md: "flex" }, mr: 1 }} />
          <Typography
            variant="h6"
            noWrap
            component="a"
            href="/home"
            sx={{
              mr: 2,
              display: { xs: "none", md: "flex" },
              fontFamily: "monospace",
              fontWeight: 700,
              letterSpacing: ".3rem",
              color: "inherit",
              textDecoration: "none",
            }}
          >
            LOGO
          </Typography>

          <Box sx={{ flexGrow: 1, display: { xs: "flex", md: "none" } }}>
            <IconButton
              size="large"
              aria-label="account of current user"
              aria-controls="menu-appbar"
              aria-haspopup="true"
              onClick={handleOpenNavMenu}
              color="inherit"
            >
              <MenuIcon />
            </IconButton>
            <Menu
              id="menu-appbar"
              anchorEl={anchorElNav}
              anchorOrigin={{
                vertical: "bottom",
                horizontal: "left",
              }}
              keepMounted
              transformOrigin={{
                vertical: "top",
                horizontal: "left",
              }}
              open={Boolean(anchorElNav)}
              onClose={handleCloseNavMenu}
              sx={{
                display: { xs: "block", md: "none" },
              }}
            >
              {pages.map((page, index) => (
                <MenuItem key={index} onClick={handleCloseNavMenu}>
                  <Typography href={page.to} component="a" textAlign="center">
                    {page.name}
                  </Typography>
                </MenuItem>
              ))}
            </Menu>
          </Box>
          <AdbIcon sx={{ display: { xs: "flex", md: "none" }, mr: 1 }} />
          <Typography
            variant="h5"
            noWrap
            component="a"
            href="/"
            sx={{
              mr: 2,
              display: { xs: "flex", md: "none" },
              flexGrow: 1,
              fontFamily: "monospace",
              fontWeight: 700,
              letterSpacing: ".3rem",
              color: "inherit",
              textDecoration: "none",
            }}
          >
            LOGO
          </Typography>
          <Box sx={{ flexGrow: 1, display: { xs: "none", md: "flex" } }}>
            {pages.map((page, index) => (
              <Button
                key={index}
                onClick={handleCloseNavMenu}
                href={page.to}
                sx={{ my: 2, color: "white", display: "block" }}
              >
                {page.name}
              </Button>
            ))}
          </Box>

          <Box sx={{ flexGrow: 0 }}>
            <Tooltip title="Open settings">
              <IconButton onClick={handleOpenUserMenu} sx={{ p: 0 }}>
                <Avatar
                  alt="Remy Sharp"
                  src="https://s3-alpha.figma.com/hub/file/1913095808/a7bdc469-cd70-4ea1-bb57-b59204ad8182-cover.png"
                />
              </IconButton>
            </Tooltip>
            <Menu
              sx={{ mt: "45px" }}
              id="menu-appbar"
              anchorEl={anchorElUser}
              anchorOrigin={{
                vertical: "top",
                horizontal: "right",
              }}
              keepMounted
              transformOrigin={{
                vertical: "top",
                horizontal: "right",
              }}
              open={Boolean(anchorElUser)}
              onClose={handleCloseUserMenu}
            >
              {settings().map((setting, index) => (
                <MenuItem key={index} onClick={handleCloseUserMenu}>
                  {setting.type === "Button" ? (
                    <Typography
                      textAlign="center"
                      component="button"
                      style={{
                        background: "none",
                        border: "none",
                      }}
                      onClick={setting.func}
                    >
                      {setting.name}
                    </Typography>
                  ) : (
                    <Typography
                      textAlign="center"
                      component="a"
                      href={setting.to}
                    >
                      {setting.name}
                    </Typography>
                  )}
                </MenuItem>
              ))}
            </Menu>
          </Box>
        </Toolbar>
      </Container>
      <ToastContainer />
    </AppBar>
  );
}
