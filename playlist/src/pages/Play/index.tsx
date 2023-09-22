import { useState, useEffect, useRef } from "react";
import YouTube from "react-youtube";
import PlayCircleIcon from "@mui/icons-material/PlayCircle";
import Button from "@mui/material/Button";
import Container from "@mui/material/Container";
import ListConnection from "../../class/connection/list";
import { useNavigate, useParams } from "react-router-dom";
import Storage from "../../class/storage";
import { User, WID } from "../../types";
import { ToastContainer, toast } from "react-toastify";
import RestartAltIcon from "@mui/icons-material/RestartAlt";

async function getCode(id: string, setVideoIds: (code: string) => void) {
  const user = Storage.local.getSession() as WID<User> | {};
  if ("id" in user) {
    const dataCodeNext = await ListConnection.getNext(id, user.id);
    if ("error" in dataCodeNext) {
      toast.error(dataCodeNext.error);
    } else if ("message" in dataCodeNext) {
      toast.success(dataCodeNext.message);
    } else {
      if (Array.isArray(dataCodeNext) && dataCodeNext.length > 0) {
        const codeNext = dataCodeNext[0];
        setVideoIds(codeNext.code);
        toast.success("Get Code Successful");
      }
    }
  }
}

function VideoPlayer() {
  const [videoId, setVideoIds] = useState<string>("");
  const { id } = useParams();
  const playerRef = useRef<YouTube | null>(null);
  const navigate = useNavigate();

  if (id === undefined) return;
  const onVideoEnd = () => {
    //TODO: Solicitud en la api par el siguiente codigo
    getCode(id, setVideoIds);
  };

  useEffect(() => {
    getCode(id, setVideoIds);
  }, []);

  useEffect(() => {
    // Esta función se ejecuta cuando cambia el currentVideoIndex
    // Cambiar el video actual cuando se cambie el código
    if (playerRef.current) {
      playerRef.current.internalPlayer.loadVideoById(videoId);
    }
    setTimeout(() => {
      if (playerRef.current) {
        playerRef.current.internalPlayer.playVideo();
      }
    }, 2000);
  }, [videoId]);

  const opts = {
    height: "390",
    width: "640",
  };

  return (
    <Container component="main">
      <Button
        variant="outlined"
        onClick={() => {
          if (playerRef.current) {
            playerRef.current.internalPlayer.playVideo();
          }
        }}
      >
        <PlayCircleIcon />
      </Button>
      <Button
        variant="outlined"
        onClick={() => {
          async function resets(id: string) {
            const reset = await ListConnection.resetById(id);
            if ("error" in reset) {
              toast.error(reset.error);
            } else if ("message" in reset) {
              toast.success(reset.message);
              navigate(0);
            }
          }
          resets(id);
        }}
      >
        <RestartAltIcon />
      </Button>
      <YouTube
        videoId={videoId}
        opts={opts}
        onEnd={onVideoEnd}
        ref={playerRef}
      />
      <ToastContainer />
    </Container>
  );
}

export default VideoPlayer;
