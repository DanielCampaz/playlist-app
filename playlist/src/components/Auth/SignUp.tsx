import Avatar from "@mui/material/Avatar";
import Button from "@mui/material/Button";
import CssBaseline from "@mui/material/CssBaseline";
import TextField from "@mui/material/TextField";
import { Link } from "react-router-dom";
import Grid from "@mui/material/Grid";
import Box from "@mui/material/Box";
import FaceIcon from "@mui/icons-material/Face";
import Typography from "@mui/material/Typography";
import Container from "@mui/material/Container";
import { createTheme, ThemeProvider } from "@mui/material/styles";
import { Copyright } from "../../const";
import { useForm } from "react-hook-form";
import { useNavigate } from "react-router-dom";
import AuthConnection from "../../class/connection/auth";
import { toast, ToastContainer } from "react-toastify";
import { User } from "../../types";

const dataSingUp = [
  {
    id: "name",
    label: "Name",
    name: "name",
    placeholder: "Jose Luis",
    type: "text",
  },
  {
    id: "lastname",
    label: "Last Name",
    name: "lastname",
    placeholder: "Rodriguez",
    type: "text",
  },
  {
    id: "email",
    label: "Email",
    name: "email",
    placeholder: "example@example.com",
    type: "email",
  },
  {
    id: "password",
    label: "Password",
    name: "password",
    placeholder: "23132654",
    type: "password",
  },
];

// TODO remove, this demo shouldn't need to reset the theme.
const defaultTheme = createTheme();

export default function SignUp() {
  const { register, handleSubmit } = useForm();
  const navigate = useNavigate();

  const handleSubmitData = async (data: any) => {
    //nsole.log(data);

    const responSingUp = await AuthConnection.SingUp(data as User);
    if (responSingUp.error) {
      toast.error(responSingUp.error);
    } else {
      toast.success("Singup Succesful");
      setTimeout(() => {
        navigate("/auth/login", {
          replace: true,
        });
      }, 6000);
    }
  };

  return (
    <ThemeProvider theme={defaultTheme}>
      <Container component="main" maxWidth="xs">
        <CssBaseline />
        <Box
          sx={{
            marginTop: 8,
            display: "flex",
            flexDirection: "column",
            alignItems: "center",
          }}
        >
          <Avatar sx={{ m: 1, bgcolor: "secondary.main" }}>
            <FaceIcon />
          </Avatar>
          <Typography component="h1" variant="h5">
            Sign up
          </Typography>
          <Box
            component="form"
            noValidate
            onSubmit={handleSubmit(handleSubmitData)}
            sx={{ mt: 3 }}
          >
            <Grid container spacing={2}>
              {dataSingUp.map((data, index) => {
                return data.name === "name" || data.name === "lastname" ? (
                  <Grid item xs={12} sm={6} key={index}>
                    <TextField
                      autoComplete={data.name}
                      required
                      fullWidth
                      id={data.id}
                      label={data.label}
                      type={data.type}
                      autoFocus
                      {...register(data.name)}
                    />
                  </Grid>
                ) : (
                  <Grid item xs={12} key={index}>
                    <TextField
                      required
                      fullWidth
                      id={data.id}
                      label={data.label}
                      type={data.type}
                      autoComplete={data.name}
                      {...register(data.name)}
                    />
                  </Grid>
                );
              })}
            </Grid>
            <Button
              type="submit"
              fullWidth
              variant="contained"
              sx={{ mt: 3, mb: 2 }}
            >
              Sign Up
            </Button>
            <Grid container justifyContent="flex-end">
              <Grid item>
                <Link to="/auth/login">Already have an account? Sign in</Link>
              </Grid>
            </Grid>
          </Box>
        </Box>
        <Copyright sx={{ mt: 5 }} />
      </Container>
      <ToastContainer />
    </ThemeProvider>
  );
}
