import React, { useMemo } from "react";

import Box from "@mui/material/Box";
import Typography from "@mui/material/Typography";

import { VideoData } from "../services";

const Song: React.FC<VideoData> = (videoData: VideoData) => {
  const img = useMemo(
    () => videoData.thumbnails[0]?.url ?? "",
    [videoData.thumbnails]
  );

  return (
    <>
      <Box
        sx={{
          display: "flex",
          margin: "0 auto 30px",
          alignItems: "center",
          position: "relative",
          overflow: "hidden",
        }}
      >
        <Box
          sx={{
            borderRadius: "50%",
            fontSize: 0,
            overflow: "hidden",
            animation: "circle 8s linear infinite",
          }}
        >
          <img
            src={img}
            alt="avatar"
            style={{
              width: "320px",
              aspectRatio: "1 / 1",
            }}
          />
          <Box
            sx={{
              width: "60px",
              height: "60px",
              background: "#ffffff",
              position: "absolute",
              top: "calc(50% - 30px)",
              left: "calc(50% - 30px)",
              borderRadius: "50%",
              boxShadow: "0 0 8px inset rgb(0 0 0 / 20%)",
            }}
          ></Box>
        </Box>
      </Box>

      <Box
        sx={{
          margin: "20px 30px",
          height: 95,
        }}
      >
        <Typography component="h2">{videoData.label}</Typography>
        <Typography component="p">
          <br />
          duration: {videoData.durationLabel}
        </Typography>
      </Box>
    </>
  );
};

export default Song;
