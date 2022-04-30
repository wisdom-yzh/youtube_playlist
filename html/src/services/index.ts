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
