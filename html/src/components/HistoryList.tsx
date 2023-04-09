import React from "react";

import Box from "@mui/material/Box";
import Typography from "@mui/material/Typography";
import { HistoryPlayList, PlayListData } from "../services";
import Stack from "@mui/material/Stack";
import Paper from "@mui/material/Paper";
import Divider from "@mui/material/Divider";
import { styled } from "@mui/material/styles";
import { useNavigate } from "react-router-dom";

const Item = styled(Paper)(({ theme }) => ({
  backgroundColor: "#eee",
  ...theme.typography.body2,
  padding: theme.spacing(2),
  color: theme.palette.text.secondary,
}));

const getTotalDuration = (item: PlayListData): string => {
  let totalSeconds = item.list.reduce((total, song) => {
    total += song.durationSeconds;
    return total;
  }, 0);

  let output = "";
  if (totalSeconds > 3600) {
    output += `${Math.floor(totalSeconds / 3600)}h`;
  }
  totalSeconds %= 3600;
  if (totalSeconds > 60) {
    output += `${Math.floor(totalSeconds / 60)}m`;
  }
  totalSeconds %= 60;
  if (totalSeconds > 0) {
    output += `${totalSeconds}s`;
  }

  return output;
};

const getPlayListID = (path: string): string => {
  return path.replace("/api/list", "/play");
};

const HistoryList: React.FC<{
  historyList: HistoryPlayList[];
}> = ({ historyList }) => {
  const navigate = useNavigate();

  if (historyList.length === 0) {
    return null;
  }

  return (
    <Box sx={{ width: "100%" }} mt={5}>
      <Typography component="p" textAlign="center">
        History
      </Typography>
      <Stack mt={2} spacing={2}>
        {historyList.map((playlist) => (
          <Item
            elevation={0}
            key={playlist.path}
            onClick={() => navigate(getPlayListID(playlist.path))}
          >
            <Stack
              direction="row"
              divider={<Divider orientation="vertical" flexItem />}
            >
              <Typography component="p" flex={2}>
                {playlist.title}
              </Typography>
              <Typography
                component="p"
                flex={1}
                paddingX={3}
                textAlign="center"
              >
                {playlist.list.length} Songs
              </Typography>
              <Typography
                component="p"
                flex={1}
                paddingX={3}
                textAlign="center"
              >
                Duration {getTotalDuration(playlist)}
              </Typography>
            </Stack>
          </Item>
        ))}
      </Stack>
    </Box>
  );
};

export default HistoryList;
