import { useState, useEffect, useRef } from "react";
import YouTube from "react-youtube";
import Container from "@mui/material/Container";
import ListConnection from "../../class/connection/list";
import { Code, Paginate } from "../../types";
import { Alert } from "@mui/material";
interface Props {
  id: string | number;
}
function TabletAdd({ id }: Props) {
  const [codes, setCodes] = useState<Code[]>([]);
  const playerRef = useRef<YouTube | null>(null);

  useEffect(() => {
    async function getData() {
      const data = (await ListConnection.getAllCodes(id)) as Paginate<Code[]>;
      console.log(data);
      setCodes(data.data);
    }
    getData();
  }, []);

  const opts = {
    height: "190",
    width: "440",
  };

  if (codes.length <= 0) {
    return <>No Codes</>;
  }

  return (
    <Container component="main">
      {codes.map((code, index) => {
        return (
          <div
            key={`playlist-code-${index}-${code.code}`}
            style={{
              margin: "10px",
            }}
          >
            <YouTube videoId={code.code} opts={opts} ref={playerRef} />
            {code.isPlatey ? (
              <Alert severity="success">
                The video with code {code.code} has already been played
              </Alert>
            ) : (
              <Alert severity="info">
                The video with code {code.code} has not been played
              </Alert>
            )}
          </div>
        );
      })}
    </Container>
  );
}

export default TabletAdd;
