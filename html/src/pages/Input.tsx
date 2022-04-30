import React, { useState } from "react";
import { useNavigate } from "react-router-dom";

import Typography from "@mui/material/Typography";
import TextField from "@mui/material/TextField";
import Box from "@mui/material/Box";
import Button from "@mui/material/Button";

const Input: React.FC = () => {
  const navigate = useNavigate();
  const [playlistId, setPlaylistId] = useState("");
  return (
    <>
      <Typography component="h1" variant="h5">
        👋 Input Youtube Playlist Here
      </Typography>
      <Box
        sx={{
          display: "flex",
          flexDirection: "row",
          width: "100%",
          marginTop: 5,
        }}
      >
        <TextField
          margin="normal"
          fullWidth
          autoFocus
          onChange={(e) => setPlaylistId(e.target.value)}
        />
        <Button
          type="submit"
          fullWidth
          variant="contained"
          sx={{ mt: 2, mb: 1, width: 64 }}
          onClick={() => navigate(`/play/${playlistId}`)}
        >
          Go
        </Button>
      </Box>
    </>
  );
};

export default Input;
