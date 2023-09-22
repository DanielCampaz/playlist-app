import Container from "@mui/material/Container";
import List from "../../components/List";
import Button from "@mui/material/Button";
import AddCircleIcon from "@mui/icons-material/AddCircle";
// import { useState } from "react";
// import ListConnection from "../../class/connection/list";

export default function Home() {
  return (
    <div
      style={{
        marginTop: "15px",
      }}
    >
      <Container component="main">
        <Button variant="outlined" href="/create">
          <AddCircleIcon />
        </Button>
        <List />
      </Container>
    </div>
  );
}
