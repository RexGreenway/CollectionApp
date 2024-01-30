import Bubble from "./Bubble";

export class CollectionNode {
  id: number;
  name: string;
  items: object[];

  constructor(id: number, name: string, items: object[]) {
    this.id = id;
    this.name = name;
    this.items = items;
  }

  radius() {
    return this.items.length;
  }
}

export default function Collection() {
  // Get Collection Data (temp data here but API call later)
  const collections: CollectionNode[] = [
    new CollectionNode(0, "Films", [{ name: "film-0" }, { name: "film-1" }]),
    new CollectionNode(1, "Comics", [{ name: "comic-0" }]),
    new CollectionNode(2, "Generic", [
      { name: "gen-0" },
      { name: "gen-1" },
      { name: "gen-2" },
    ]),
    new CollectionNode(3, "Stamps", [{ name: "stamp-0" }, { name: "stamp-1" }]),
    new CollectionNode(4, "Coins", [
      { name: "coin-0" },
      { name: "coin-1" },
      { name: "coin-2" },
    ]),
    new CollectionNode(5, "Beanie Babies", [
      { name: "bb-0" },
      { name: "bb-1" },
    ]),
    new CollectionNode(2, "Generic", [
      { name: "gen-0" },
      { name: "gen-1" },
      { name: "gen-2" },
    ]),
    new CollectionNode(3, "Stamps", [{ name: "stamp-0" }, { name: "stamp-1" }]),
    new CollectionNode(4, "Coins", [
      { name: "coin-0" },
      { name: "coin-1" },
      { name: "coin-2" },
    ]),
    new CollectionNode(5, "Beanie Babies", [
      { name: "bb-0" },
      { name: "bb-1" },
    ]),
    new CollectionNode(2, "Generic", [
      { name: "gen-0" },
      { name: "gen-1" },
      { name: "gen-2" },
    ]),
    new CollectionNode(3, "Stamps", [{ name: "stamp-0" }, { name: "stamp-1" }]),
    new CollectionNode(4, "Coins", [
      { name: "coin-0" },
      { name: "coin-1" },
      { name: "coin-2" },
    ]),
    new CollectionNode(5, "Beanie Babies", [
      { name: "bb-0" },
      { name: "bb-1" },
    ]),
  ];

  return <Bubble>{collections}</Bubble>;
}
