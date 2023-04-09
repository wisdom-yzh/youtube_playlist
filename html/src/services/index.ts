import { useEffect, useState } from "react";
import { useService, fetcher } from "./common";

export interface Thumbnails {
  url: string;
  width: number;
  height: number;
}

export interface VideoData {
  thumbnails: Thumbnails[];
  vid: string;
  name: string;
  label: string;
  durationSeconds: number;
  duration: string;
  durationLabel: string;
}

export interface PlayListData {
  title: string;
  list: VideoData[];
}

export const usePlayList = (playlistId: string) => {
  return useService<PlayListData>(`/api/list/${playlistId}`);
};

export interface VideoUrl {
  url: string;
}

export const fetchVideoUrl = (videoId: string) => {
  return fetcher<VideoUrl>(`/api/video/${videoId}`);
};

export interface HistoryPlayList extends PlayListData {
  path: string;
}

export const useHistoryPlayList = () => {
  const [historyPlayList, setHistoryPlayList] = useState<HistoryPlayList[]>([]);

  useEffect(() => {
    (async () => {
      const cache = await caches.open("player-list");
      const keys = await cache.keys();
      if (keys.length === historyPlayList.length) {
        return;
      }

      const playlist = [];
      for (const key of keys) {
        const res = await cache.match(key);
        if (res === undefined) {
          continue;
        }
        const data = (await res?.json()) as HistoryPlayList;
        data.path = new URL(key.url).pathname;

        playlist.push(data);
      }

      setHistoryPlayList(playlist);
    })();
  }, [historyPlayList.length]);

  return historyPlayList;
};
