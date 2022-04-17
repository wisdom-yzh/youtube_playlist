import React, { useCallback, useState } from "react";
import "./App.css";
import { usePlayList, fetchVideoUrl } from "./service";

const DEFAULT_PLAYLIST = "PL2wrxGo0Q1AX8QMpUwUwGWpzoZqkObUNT";

const App: React.FC = () => {
  const url = new URL(window.location.href);
  const { data, error } = usePlayList(
    url.searchParams.get("list") ?? DEFAULT_PLAYLIST
  );
  const [urls, setUrls] = useState<Record<string, string>>({});
  const gotoVideo = useCallback(
    (id: string) => {
      if (urls[id]) {
        window.location.href = urls[id];
        return;
      }
      fetchVideoUrl(id).then((data) => {
        setUrls({
          ...urls,
          id: data.url,
        });
        window.location.href = data.url;
      });
    },
    [urls, setUrls]
  );

  return (
    <div className="App">
      {error !== undefined ? <div>failed to parse playlist</div> : null}
      <h3>{data?.title}</h3>
      <ul>
        {data?.list.map((video) => (
          <li>
            <p>Name: {video.name}</p>
            <p>Label: {video.label}</p>
            <p>Length{video.durationLabel}</p>
            <p>
              <button onClick={() => gotoVideo(video.vid)}>Listen</button>
            </p>
          </li>
        ))}
      </ul>
    </div>
  );
};

export default App;
