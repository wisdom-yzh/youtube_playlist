import useSWR from "swr";

const hostname = process.env.REACT_APP_API_SERVICE ?? "";

export const fetcher = async <T>(url: string): Promise<T> => {
  return fetch(`${hostname}${url}`, { mode: "cors" }).then(
    (res) => res.json() as unknown as T
  );
};

export const useService = <T>(path: string) => {
  return useSWR<T>(path, fetcher, {
    refreshInterval: 0,
    revalidateIfStale: false,
    revalidateOnFocus: false,
  });
};
