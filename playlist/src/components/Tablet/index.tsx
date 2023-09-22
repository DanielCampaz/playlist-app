import { alpha } from "@mui/material/styles";
import Box from "@mui/material/Box";
import Table from "@mui/material/Table";
import TableBody from "@mui/material/TableBody";
import TableCell from "@mui/material/TableCell";
import TableContainer from "@mui/material/TableContainer";
import TableHead from "@mui/material/TableHead";
import TablePagination from "@mui/material/TablePagination";
import TableRow from "@mui/material/TableRow";
import TableSortLabel from "@mui/material/TableSortLabel";
import Toolbar from "@mui/material/Toolbar";
import Typography from "@mui/material/Typography";
import Paper from "@mui/material/Paper";
import Checkbox from "@mui/material/Checkbox";
import IconButton from "@mui/material/IconButton";
import Tooltip from "@mui/material/Tooltip";
import DeleteIcon from "@mui/icons-material/Delete";
import FilterListIcon from "@mui/icons-material/FilterList";
import { visuallyHidden } from "@mui/utils";
import { useMemo, useState, MouseEvent } from "react";
import PlayCircleIcon from "@mui/icons-material/PlayCircle";
import AddCircleIcon from "@mui/icons-material/AddCircle";
import {
  GetProperty,
  GetPropertyGuide,
  Order,
  TableCellSort,
  getComparator,
  stableSort,
} from "./utilsTable";
import { urlAdd, urlPlay } from "../../const";
import ListConnection from "../../class/connection/list";
import { ToastContainer, toast } from "react-toastify";

// interface Data {
//   calories: number;
//   carbs: number;
//   fat: number;
//   name: string;
//   protein: number;
// }

// function createData(
//   name: string,
//   calories: number,
//   fat: number,
//   carbs: number,
//   protein: number
// ): Data {
//   return {
//     name,
//     calories,
//     fat,
//     carbs,
//     protein,
//   };
// }

// const rows = [
//   createData("Cupcake", 305, 3.7, 67, 4.3),
//   createData("Donut", 452, 25.0, 51, 4.9),
//   createData("Eclair", 262, 16.0, 24, 6.0),
//   createData("Frozen yoghurt", 159, 6.0, 24, 4.0),
//   createData("Gingerbread", 356, 16.0, 49, 3.9),
//   createData("Honeycomb", 408, 3.2, 87, 6.5),
//   createData("Ice cream sandwich", 237, 9.0, 37, 4.3),
//   createData("Jelly Bean", 375, 0.0, 94, 0.0),
//   createData("KitKat", 518, 26.0, 65, 7.0),
//   createData("Lollipop", 392, 0.2, 98, 0.0),
//   createData("Marshmallow", 318, 0, 81, 2.0),
//   createData("Nougat", 360, 19.0, 9, 37.0),
//   createData("Oreo", 437, 18.0, 63, 4.0),
// ];

// interface HeadCell {
//   disablePadding: boolean;
//   id: keyof Data;
//   label: string;
//   numeric: boolean;
// }

// const headCells: readonly HeadCell[] = [
//   {
//     id: "name",
//     numeric: false,
//     disablePadding: true,
//     label: "Dessert (100g serving) 65465465465465",
//   },
//   {
//     id: "calories",
//     numeric: true,
//     disablePadding: false,
//     label: "Calories",
//   },
//   {
//     id: "fat",
//     numeric: true,
//     disablePadding: false,
//     label: "Fat (g)",
//   },
//   {
//     id: "carbs",
//     numeric: true,
//     disablePadding: false,
//     label: "Carbs (g)",
//   },
//   {
//     id: "protein",
//     numeric: true,
//     disablePadding: false,
//     label: "Protein (g)",
//   },
// ];

interface EnhancedTableProps<T extends { [key: string]: any }> {
  numSelected: number;
  onRequestSort: (event: React.MouseEvent<unknown>, property: keyof T) => void;
  onSelectAllClick: (event: React.ChangeEvent<HTMLInputElement>) => void;
  order: Order;
  data: T[];
  orderBy: string | number | symbol;
  rowCount: number;
}

function EnhancedTableHead<T extends { [key: string]: any }>(
  props: EnhancedTableProps<T>
) {
  const {
    onSelectAllClick,
    order,
    orderBy,
    numSelected,
    rowCount,
    data,
    onRequestSort,
  } = props;
  const dataSorted = TableCellSort(data);
  const createSortHandler =
    (property: keyof T) => (event: React.MouseEvent<unknown>) => {
      onRequestSort(event, property);
    };

  return (
    <TableHead>
      <TableRow>
        <TableCell padding="checkbox">
          <Checkbox
            color="primary"
            indeterminate={numSelected > 0 && numSelected < rowCount}
            checked={rowCount > 0 && numSelected === rowCount}
            onChange={onSelectAllClick}
            inputProps={{
              "aria-label": "select all desserts",
            }}
          />
        </TableCell>

        {dataSorted.map((headCell) => (
          <TableCell
            key={headCell.id}
            align={headCell.numeric ? "right" : "left"}
            padding={headCell.disablePadding ? "none" : "normal"}
            sortDirection={orderBy === headCell.id ? order : false}
          >
            <TableSortLabel
              active={orderBy === headCell.id}
              direction={orderBy === headCell.id ? order : "asc"}
              onClick={createSortHandler(headCell.id)}
            >
              {headCell.label}
              {orderBy === headCell.id ? (
                <Box component="span" sx={visuallyHidden}>
                  {order === "desc" ? "sorted descending" : "sorted ascending"}
                </Box>
              ) : null}
            </TableSortLabel>
          </TableCell>
        ))}
        <TableCell
          key={"play"}
          align={"right"}
          padding={"normal"}
          sortDirection={orderBy === "play" ? order : false}
        >
          <TableSortLabel
            active={orderBy === "play"}
            direction={orderBy === "play" ? order : "asc"}
            onClick={createSortHandler("play")}
          >
            {"Play"}
            {orderBy === "play" ? (
              <Box component="span" sx={visuallyHidden}>
                {order === "desc" ? "sorted descending" : "sorted ascending"}
              </Box>
            ) : null}
          </TableSortLabel>
        </TableCell>
        <TableCell
          key={"add"}
          align={"right"}
          padding={"normal"}
          sortDirection={orderBy === "add" ? order : false}
        >
          <TableSortLabel
            active={orderBy === "add"}
            direction={orderBy === "add" ? order : "asc"}
            onClick={createSortHandler("add")}
          >
            {"add"}
            {orderBy === "add" ? (
              <Box component="span" sx={visuallyHidden}>
                {order === "desc" ? "sorted descending" : "sorted ascending"}
              </Box>
            ) : null}
          </TableSortLabel>
        </TableCell>
      </TableRow>
    </TableHead>
  );
}

interface EnhancedTableToolbarProps {
  selected: (string | number)[];
  title: string;
}

function EnhancedTableToolbar({ selected, title }: EnhancedTableToolbarProps) {
  const numSelected = selected.length;
  return (
    <Toolbar
      sx={{
        pl: { sm: 2 },
        pr: { xs: 1, sm: 1 },
        ...(numSelected > 0 && {
          bgcolor: (theme) =>
            alpha(
              theme.palette.primary.main,
              theme.palette.action.activatedOpacity
            ),
        }),
      }}
    >
      {numSelected > 0 ? (
        <Typography
          sx={{ flex: "1 1 100%" }}
          color="inherit"
          variant="subtitle1"
          component="div"
        >
          {numSelected} selected
        </Typography>
      ) : (
        <Typography
          sx={{ flex: "1 1 100%" }}
          variant="h6"
          id="tableTitle"
          component="div"
        >
          {title}
        </Typography>
      )}
      {numSelected > 0 ? (
        <Tooltip title="Delete">
          <IconButton
            onClick={async () => {
              // FIXME: DELTE MANYYYYYYYYYYYYYYY
              const dataer = await ListConnection.deleteManyTable(selected);
              console.log({
                dataer,
                selected,
              });
              toast.success("Codes Delete");
            }}
          >
            <DeleteIcon />
          </IconButton>
        </Tooltip>
      ) : (
        <Tooltip title="Filter list">
          <IconButton>
            <FilterListIcon />
          </IconButton>
        </Tooltip>
      )}
    </Toolbar>
  );
}

interface TableProps<T> {
  data: T[];
  title: string;
}

export default function Tableet<T extends { [key: string]: any }>({
  data,
  title,
}: TableProps<T>) {
  const [order, setOrder] = useState<Order>("asc");
  const [orderBy, setOrderBy] = useState<keyof T>("");
  const [selected, setSelected] = useState<(string | number)[]>([]);
  const [page, setPage] = useState(0);
  const [rowsPerPage, setRowsPerPage] = useState(5);

  const handleRequestSort = (
    _event: MouseEvent<unknown>,
    property: keyof T
  ) => {
    const isAsc = orderBy === property && order === "asc";
    setOrder(isAsc ? "desc" : "asc");
    setOrderBy(property);
  };

  const handleSelectAllClick = (event: React.ChangeEvent<HTMLInputElement>) => {
    if (event.target.checked) {
      const newSelected = data.map((n) => n[GetPropertyGuide(n)]);
      setSelected(newSelected);
      return;
    }
    setSelected([]);
  };

  const handleClick = (_event: MouseEvent<unknown>, name: string | number) => {
    const selectedIndex = selected.indexOf(name);
    let newSelected: (string | number)[] = [];

    if (selectedIndex === -1) {
      newSelected = newSelected.concat(selected, name);
    } else if (selectedIndex === 0) {
      newSelected = newSelected.concat(selected.slice(1));
    } else if (selectedIndex === selected.length - 1) {
      newSelected = newSelected.concat(selected.slice(0, -1));
    } else if (selectedIndex > 0) {
      newSelected = newSelected.concat(
        selected.slice(0, selectedIndex),
        selected.slice(selectedIndex + 1)
      );
    }

    setSelected(newSelected);
  };

  const handleChangePage = (_event: unknown, newPage: number) => {
    setPage(newPage);
  };

  const handleChangeRowsPerPage = (
    event: React.ChangeEvent<HTMLInputElement>
  ) => {
    setRowsPerPage(parseInt(event.target.value, 10));
    setPage(0);
  };

  const isSelected = (name: string | number) => selected.indexOf(name) !== -1;

  // Avoid a layout jump when reaching the last page with empty rows.
  const emptyRows =
    page > 0 ? Math.max(0, (1 + page) * rowsPerPage - data.length) : 0;

  const visibleRows = useMemo(
    () =>
      stableSort(data, getComparator(order, orderBy)).slice(
        page * rowsPerPage,
        page * rowsPerPage + rowsPerPage
      ),
    [order, orderBy, page, rowsPerPage]
  );
  return (
    <Box sx={{ width: "100%" }}>
      <Paper sx={{ width: "100%", mb: 2 }}>
        <EnhancedTableToolbar title={title} selected={selected} />
        <TableContainer>
          <Table
            sx={{ minWidth: 750 }}
            aria-labelledby="tableTitle"
            size={"medium"}
          >
            <EnhancedTableHead
              numSelected={selected.length}
              order={order}
              data={data}
              orderBy={orderBy}
              onSelectAllClick={handleSelectAllClick}
              onRequestSort={handleRequestSort}
              rowCount={data.length}
            />
            <TableBody>
              {visibleRows.map((row, index) => {
                const propertyGuide = row[GetPropertyGuide(row)];
                const isItemSelected = isSelected(propertyGuide);
                const labelId = `enhanced-table-checkbox-${index}-${
                  row[GetPropertyGuide(row)]
                }`;
                const keys = Object.keys(row);

                return (
                  <TableRow
                    hover
                    onClick={(event) => handleClick(event, propertyGuide)}
                    role="checkbox"
                    aria-checked={isItemSelected}
                    tabIndex={-1}
                    key={labelId}
                    selected={isItemSelected}
                    sx={{ cursor: "pointer" }}
                  >
                    <TableCell padding="checkbox">
                      <Checkbox
                        color="primary"
                        checked={isItemSelected}
                        inputProps={{
                          "aria-labelledby": labelId,
                        }}
                      />
                    </TableCell>
                    {keys.map((key, indexe) => {
                      const tableRowId = `enhanced-table-checkbox-${indexe}`;

                      row[key];
                      if (indexe === 0) {
                        return (
                          <TableCell
                            component="th"
                            id={labelId}
                            scope="row"
                            padding="none"
                            key={tableRowId + "645"}
                          >
                            {row[key]}
                          </TableCell>
                        );
                      } else {
                        return (
                          <TableCell align="right" key={tableRowId + "6452"}>
                            {" "}
                            {row[key]}
                          </TableCell>
                        );
                      }
                    })}
                    <TableCell align="right" key={labelId + "38223"}>
                      <Typography
                        textAlign="center"
                        component="a"
                        href={urlPlay(GetProperty(row))}
                      >
                        <PlayCircleIcon />
                      </Typography>
                    </TableCell>
                    <TableCell align="right" key={labelId + "3223"}>
                      <Typography
                        textAlign="center"
                        component="a"
                        href={urlAdd(GetProperty(row))}
                      >
                        <AddCircleIcon />
                      </Typography>
                    </TableCell>
                  </TableRow>
                );
              })}
              {emptyRows > 0 && (
                <TableRow
                  style={{
                    height: 53 * emptyRows,
                  }}
                >
                  <TableCell colSpan={6} />
                </TableRow>
              )}
            </TableBody>
          </Table>
        </TableContainer>
        <TablePagination
          rowsPerPageOptions={[5, 10, 25]}
          component="div"
          count={data.length}
          rowsPerPage={rowsPerPage}
          page={page}
          onPageChange={handleChangePage}
          onRowsPerPageChange={handleChangeRowsPerPage}
        />
      </Paper>
      <ToastContainer />
    </Box>
  );
}
