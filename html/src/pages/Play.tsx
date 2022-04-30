import React, { useState } from "react";
import { useParams } from "react-router-dom";

import Box from "@mui/material/Box";
import Alert from "@mui/material/Alert";
import AlertTitle from "@mui/material/AlertTitle";
import Typography from "@mui/material/Typography";

import { usePlayList } from "../services";
import Song from "../components/Song";
import PlayerControl from "../components/PlayerControl";

const Play: React.FC = () => {
  const { playlistId } = useParams<{ playlistId: string }>();
  const { data: playlist, error } = usePlayList(playlistId ?? "");
  const [trackIndex, setTrackIndex] = useState(0);

  return (
    <Box
      sx={{
        alignItems: "center",
        display: "flex",
        flexDirection: "column",
        width: "100%",
      }}
    >
      {error ? (
        <Alert severity="error" sx={{ width: "100%" }}>
          <AlertTitle>Error</AlertTitle>
          Failed to fetch playlist... <strong>Check the playlist id!</strong>
        </Alert>
      ) : !playlist?.list ? (
        <Alert severity="info" sx={{ width: "100%" }}>
          <AlertTitle>Info</AlertTitle>
          Fetching data...
        </Alert>
      ) : (
        <>
          <Typography component="h1" variant="h5" sx={{ mb: "30px" }}>
            {playlist.title}
          </Typography>
          <Song {...playlist.list[trackIndex]} />
          <PlayerControl
            setTrackIndex={setTrackIndex}
            trackIndex={trackIndex}
            videos={playlist.list}
          />
        </>
      )}
    </Box>
  );
};

export default Play;
