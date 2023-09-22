import { Button, Container, MenuItem, TextField } from "@mui/material";
import { useState } from "react";
import { useForm } from "react-hook-form";
import { useParams } from "react-router-dom";
import ListConnection from "../../class/connection/list";
import Storage from "../../class/storage";
import { AddType, User, WID } from "../../types";
import { ToastContainer, toast } from "react-toastify";
import TabletAdd from "../../components/TabletAdd";

const select = [
  {
    value: "code",
    label: "Code",
  },
  {
    value: "iframe",
    label: "IFrame",
  },
  {
    value: "url",
    label: "Url",
  },
];

export default function Add() {
  const { id } = useParams();
  const { register, handleSubmit } = useForm();
  const [code, setCode] = useState<string>("code");
  if (id === undefined) return;

  const handleSubmitData = async (data: any) => {
    const addtype: AddType = {
      ...data,
    };
    const session = Storage.local.getSession() as WID<User>;
    const addcomplete = await ListConnection.postAdd(id, session.id, addtype);
    if (addcomplete.error) {
      toast.error(addcomplete.error);
    } else {
      toast.success("List Create Succesful");
    }
  };
  return (
    <Container component="main" maxWidth="xs">
      <form onSubmit={handleSubmit(handleSubmitData)}>
        <TextField
          margin="normal"
          required
          fullWidth
          id="ifr"
          label="IFrame"
          autoFocus
          {...register("ifr")}
        />
        <TextField
          id="type"
          select
          label="Code"
          value={code}
          helperText="Please select your type"
          {...register("type")}
          onChange={(event) => {
            setCode(event.target.value);
          }}
        >
          {select.map((option, index) => (
            <MenuItem key={index} value={option.value}>
              {option.label}
            </MenuItem>
          ))}
        </TextField>
        <Button
          type="submit"
          fullWidth
          variant="contained"
          sx={{ mt: 3, mb: 2 }}
        >
          Add Code
        </Button>
      </form>
      <TabletAdd id={id} />
      <ToastContainer />
    </Container>
  );
}
