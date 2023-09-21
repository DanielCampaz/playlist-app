export interface TableCellReturn {
  id: string;
  numeric: boolean;
  disablePadding: boolean;
  label: string;
}
export function TableCellSort<T extends { [key: string]: any }>(
  data: Array<T>
): TableCellReturn[] {
  // Paso 1: Obtener las propiedades únicas de los objetos en data
  const uniqueProperties: string[] = [];
  data.forEach((obj) => {
    for (const prop in obj) {
      if (!uniqueProperties.includes(prop)) {
        uniqueProperties.push(prop);
      }
    }
  });

  // Paso 2: Crear la matriz tablecell basada en las propiedades únicas
  const tablecell = uniqueProperties.map((prop) => {
    return {
      id: prop.toLowerCase(),
      numeric: typeof data[0][prop] === "number",
      disablePadding: false,
      label: prop.charAt(0).toUpperCase() + prop.slice(1),
    };
  });

  return tablecell;
}

export function descendingComparator<T>(a: T, b: T, orderBy: keyof T) {
  if (b[orderBy] < a[orderBy]) {
    return -1;
  }
  if (b[orderBy] > a[orderBy]) {
    return 1;
  }
  return 0;
}

export function GetPropertyGuide<T extends { [key: string]: any }>(data: T) {
  const keys = Object.keys(data);
  if (keys.includes("id")) {
    return "id";
  }
  for (let i = 0; i < keys.length; i++) {
    const property = data[keys[i]];
    if (typeof property === "string" && property.length < 15) {
      return keys[i];
    }
  }
  return "";
}

export type Order = "asc" | "desc";

export function getComparator<Key extends keyof any>(
  order: Order,
  orderBy: Key
): (
  a: { [key in Key]: number | string },
  b: { [key in Key]: number | string }
) => number {
  return order === "desc"
    ? (a, b) => descendingComparator(a, b, orderBy)
    : (a, b) => -descendingComparator(a, b, orderBy);
}

// Since 2020 all major browsers ensure sort stability with Array.prototype.sort().
// stableSort() brings sort stability to non-modern browsers (notably IE11). If you
// only support modern browsers you can replace stableSort(exampleArray, exampleComparator)
// with exampleArray.slice().sort(exampleComparator)
export function stableSort<T>(
  array: readonly T[],
  comparator: (a: T, b: T) => number
) {
  const stabilizedThis = array.map((el, index) => [el, index] as [T, number]);
  stabilizedThis.sort((a, b) => {
    const order = comparator(a[0], b[0]);
    if (order !== 0) {
      return order;
    }
    return a[1] - b[1];
  });
  return stabilizedThis.map((el) => el[0]);
}
