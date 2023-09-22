import { Button, Container, TextField } from "@mui/material";
import { useForm } from "react-hook-form";
import Storage from "../../class/storage";
import { List, User, WID } from "../../types";
import ListConnection from "../../class/connection/list";
import { ToastContainer, toast } from "react-toastify";

export default function CreateList() {
  const { register, handleSubmit } = useForm();
  const fields = [
    {
      id: "name",
      label: "Name of Playlist",
      name: "name",
    },
  ];

  const handleSubmitData = async (data: any) => {
    const datauser = Storage.local.getSession() as WID<User>;
    const newList: Partial<List> = {
      ...data,
      iduser: datauser.id,
    };
    const listCreate = await ListConnection.postCreate(newList);
    if (listCreate.error) {
      toast.error(listCreate.error);
    } else {
      toast.success("List Create Succesful");
    }
  };
  return (
    <Container component="main" maxWidth="xs">
      <form onSubmit={handleSubmit(handleSubmitData)}>
        {fields.map((field, index) => {
          return (
            <TextField
              key={index}
              margin="normal"
              required
              fullWidth
              id={field.id}
              label={field.label}
              autoFocus
              {...register(field.name)}
            />
          );
        })}
        <Button
          type="submit"
          fullWidth
          variant="contained"
          sx={{ mt: 3, mb: 2 }}
        >
          Create Playlist
        </Button>
      </form>
      <ToastContainer />
    </Container>
  );
}
