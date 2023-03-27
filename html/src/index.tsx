import React from "react";
import ReactDOM from "react-dom/client";
import { BrowserRouter as Router, Routes, Route } from "react-router-dom";

import Input from "./pages/Input";
import Play from "./pages/Play";

import Container from "@mui/material/Container";
import Box from "@mui/material/Box";

import "./index.css";

import "@fontsource/roboto/300.css";
import "@fontsource/roboto/400.css";
import "@fontsource/roboto/500.css";
import "@fontsource/roboto/700.css";

const root = ReactDOM.createRoot(
  document.getElementById("root") as HTMLElement
);

root.render(
  <React.StrictMode>
    <Container component="main" maxWidth="md">
      <Box
        sx={{
          marginTop: 8,
          display: "flex",
          flexDirection: "column",
          alignItems: "center",
        }}
      >
        <Router>
          <Routes>
            <Route path="/*" element={<Input />} />
            <Route path="/play/:playlistId" element={<Play />} />
          </Routes>
        </Router>
      </Box>
    </Container>
  </React.StrictMode>
);
