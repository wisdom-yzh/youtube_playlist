import React, { useCallback, useEffect, useRef, useState } from "react";
import { fetchVideoUrl, VideoData } from "../services";

import Box from "@mui/material/Box";
import Stack from "@mui/material/Stack";
import Button from "@mui/material/Button";
import Typography from "@mui/material/Typography";
import LinearProgress, {
  LinearProgressProps,
} from "@mui/material/LinearProgress";
import SkipPreviousRoundedIcon from "@mui/icons-material/SkipPreviousRounded";
import SkipNextRoundedIcon from "@mui/icons-material/SkipNextRounded";
import PlayArrowRoundedIcon from "@mui/icons-material/PlayArrowRounded";
import PauseRoundedIcon from "@mui/icons-material/PauseRounded";

type PlayerControlProps = {
  setTrackIndex: (index: number) => void;
  trackIndex: number;
  videos: VideoData[];
};

const LinearProgressWithLabel = (
  props: LinearProgressProps & { value: number }
) => {
  return (
    <Box sx={{ display: "flex", alignItems: "center", marginBottom: "10px" }}>
      <Box sx={{ width: "100%", mr: 1 }}>
        <LinearProgress variant="determinate" {...props} />
      </Box>
      <Box sx={{ minWidth: 35 }}>
        <Typography variant="body2" color="text.secondary">{`${Math.round(
          props.value
        )}%`}</Typography>
      </Box>
    </Box>
  );
};

const PlayerControl: React.FC<PlayerControlProps> = ({
  setTrackIndex,
  trackIndex,
  videos,
}) => {
  const intervalRef = useRef<NodeJS.Timer | null>(null);
  const audioRef = useRef<HTMLAudioElement>(null);
  const [trackProgress, setTrackProgress] = useState(0);
  const [isPlaying, setIsPlaying] = useState(false);

  const pause = useCallback(() => {
    audioRef.current?.pause();
    setIsPlaying(false);
  }, []);

  const play = useCallback(() => {
    audioRef.current?.play().then(() => {
      setIsPlaying(true);
    });
  }, []);

  const gotoNextTrack = useCallback(() => {
    pause();
    setTrackIndex(trackIndex === videos.length - 1 ? 0 : trackIndex + 1);
  }, [videos.length, trackIndex, pause, setTrackIndex]);

  const gotoPrevTrack = useCallback(() => {
    pause();
    setTrackIndex(trackIndex === 0 ? videos.length - 1 : trackIndex - 1);
  }, [videos.length, trackIndex, pause, setTrackIndex]);

  useEffect(() => {
    const videoId = videos[trackIndex].vid;
    const url = `/api/video_data/${videoId}`;
    audioRef.current?.setAttribute("src", url);
    audioRef.current?.load();
    play();

    intervalRef.current = setInterval(() => {
      console.log("tick");
      if (audioRef.current?.ended) {
        clearInterval(intervalRef.current!);
        gotoNextTrack();
        return;
      }

      setTrackProgress(
        Math.round(
          (100 * Number(audioRef.current?.currentTime ?? 0)) /
            videos[trackIndex].durationSeconds
        )
      );
    }, 1000);

    return () => {
      if (intervalRef.current !== null) {
        clearInterval(intervalRef.current);
      }
    };
  }, [trackIndex, setTrackIndex, gotoNextTrack, play, videos]);

  return (
    <Box
      sx={{
        position: "absolute",
        bottom: 0,
        zIndex: 99,
        width: "90%",
        maxWidth: "760px",
        display: "flex",
        flexDirection: "column",
        padding: "30px 0",
        alignItems: "center",
      }}
    >
      <Box sx={{ width: "100%" }}>
        <LinearProgressWithLabel
          value={trackProgress}
          sx={{ height: "10px" }}
        />
      </Box>
      <Stack direction="row" spacing={2}>
        <Button
          variant="contained"
          endIcon={<SkipPreviousRoundedIcon />}
          onClick={gotoPrevTrack}
        >
          Prev
        </Button>
        <Button variant="contained" onClick={isPlaying ? pause : play}>
          {isPlaying ? <PauseRoundedIcon /> : <PlayArrowRoundedIcon />}
        </Button>
        <Button
          variant="contained"
          startIcon={<SkipNextRoundedIcon />}
          onClick={gotoNextTrack}
        >
          Next
        </Button>
      </Stack>
      <audio ref={audioRef} style={{ display: "none" }} />
    </Box>
  );
};

export default PlayerControl;
