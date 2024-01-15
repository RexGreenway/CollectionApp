const Item = ({ children }: { children: string }) => {
  return <h3>{children}</h3>;
};

export default function Collection() {
  const collection = [
    { ID: "1", Name: "test-film-1", Year: 1990 },
    { ID: "2", Name: "test-film-2", Year: 1999 },
    { ID: "3", Name: "test-film-3", Year: 2009 },
    { ID: "4", Name: "test-film-4", Year: 2020 },
    { ID: "5", Name: "test-film-5", Year: 3005 },
  ];

  const arr = [];
  for (const item of collection) {
    arr.push(<Item>{item.Name}</Item>);
  }

  return <div>{arr}</div>;
}
