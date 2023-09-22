import { ToastContainer, toast } from "react-toastify";
import ListConnection from "../../class/connection/list";
import Storage from "../../class/storage";
import Tablet from "../Tablet";
import { useEffect, useState } from "react";
import { Paginate } from "../../types";

// const data = [
//   {
//     id: 8,
//     userId: 1,
//     title:
//       "sunt aut facere repellat provident occaecati excepturi optio reprehenderit",
//     body: "quia et suscipit\nsuscipit recusandae consequuntur expedita et cum\nreprehenderit molestiae ut ut quas totam\nnostrum rerum est autem sunt rem eveniet architecto",
//   },
//   {
//     userId: 1,
//     id: 2,
//     title: "qui est esse",
//     body: "est rerum tempore vitae\nsequi sint nihil reprehenderit dolor beatae ea dolores neque\nfugiat blanditiis voluptate porro vel nihil molestiae ut reiciendis\nqui aperiam non debitis possimus qui neque nisi nulla",
//   },
// ];

async function getData() {
  const token = Storage.local.getToken();
  if (token === null || token === "") {
    return {
      data: [],
      title: "No Data",
    };
  } else {
    const lists = (await ListConnection.getAll()) as
      | Paginate<any>
      | { error: string };
    if ("error" in lists) {
      toast.error(lists.error);
      return {
        data: [],
        title: "Error to get Lists",
      };
    } else {
      toast.success("Get List Succesful");
    }
    // TODO: Get Lists
    return {
      data: lists.data,
      title: "Lists",
    };
  }
}

export default function List() {
  const [datas, setDatas] = useState({
    data: [],
    title: "Error to get Lists",
  });

  useEffect(() => {
    async function getDataAsync() {
      const dataass = await getData();
      //const lists = await ListConnection.getAll();
      setDatas(dataass);
    }
    getDataAsync();
  }, []);

  return (
    <div>
      <Tablet title={datas.title} data={datas.data} />
      <ToastContainer />
    </div>
  );
}
