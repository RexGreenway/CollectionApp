import * as d3 from "d3";
import { SimulationNodeDatum } from "d3";
import { useRef, useEffect, useState } from "react";

// Extends the Node type used in simulations
type Collection = {
  id: number;
  name: string;
  items: object[];
} & SimulationNodeDatum;

// Bubble defines a function creating a bubble
const Bubble = () => {
  const ref = useRef(null);

  // Get Collection Data (temp data here but API call later)
  const data: Collection[] = [
    {
      id: 0,
      name: "Films",
      items: [{ name: "film-0" }, { name: "film-1" }],
    },
    {
      id: 1,
      name: "Comics",
      items: [{ name: "comic-0" }],
    },
    {
      id: 2,
      name: "Generic",
      items: [{ name: "gen-0" }, { name: "gen-1" }, { name: "gen-2" }],
    },
    {
      id: 3,
      name: "Stamps",
      items: [{ name: "stamp-0" }, { name: "stamp-1" }],
    },

    {
      id: 4,
      name: "Coins",
      items: [{ name: "coin-0" }, { name: "coin-1" }, { name: "coin-2" }],
    },

    {
      id: 5,
      name: "Beanie Babies",
      items: [{ name: "bb-0" }, { name: "bb-1" }],
    },
    {
      id: 2,
      name: "Generic",
      items: [{ name: "gen-0" }, { name: "gen-1" }, { name: "gen-2" }],
    },
    {
      id: 3,
      name: "Stamps",
      items: [{ name: "stamp-0" }, { name: "stamp-1" }],
    },

    {
      id: 4,
      name: "Coins",
      items: [{ name: "coin-0" }, { name: "coin-1" }, { name: "coin-2" }],
    },

    {
      id: 5,
      name: "Beanie Babies",
      items: [{ name: "bb-0" }, { name: "bb-1" }],
    },
    {
      id: 2,
      name: "Generic",
      items: [{ name: "gen-0" }, { name: "gen-1" }, { name: "gen-2" }],
    },
    {
      id: 3,
      name: "Stamps",
      items: [{ name: "stamp-0" }, { name: "stamp-1" }],
    },

    {
      id: 4,
      name: "Coins",
      items: [{ name: "coin-0" }, { name: "coin-1" }, { name: "coin-2" }],
    },

    {
      id: 5,
      name: "Beanie Babies",
      items: [{ name: "bb-0" }, { name: "bb-1" }],
    },
  ];

  const [nodes, SetCol] = useState<Collection[]>(data);

  // D3 BUBBLE
  const width = 800;
  const height = 800;

  const r = 20;

  const Remove = () => {
    const temp = nodes.map((d) => ({ ...d }));
    temp.pop();
    SetCol(temp);
  };

  const Add = () => {
    const temp = nodes.map((d) => ({ ...d }));
    SetCol(
      temp.concat({
        id: 5,
        name: "Beanie Babies",
        items: [{ name: "bb-0" }, { name: "bb-1" }],
      })
    );
  };

  useEffect(() => {
    // Specify the color scale.
    const color = d3.scaleOrdinal(d3.schemeCategory10);

    // Create and start simulation
    d3.forceSimulation(nodes)
      // Centering nodes toward (0, 0) per node
      .force("x", d3.forceX().strength(0.01))
      .force("y", d3.forceY().strength(0.01))
      // Force for 'border' of nodes creates collisions
      .force(
        "collide",
        d3
          .forceCollide<Collection>((d) => r * d.items.length + 2)
          .iterations(12)
      )
      // Inter-node gravity
      .force(
        "charge",
        d3.forceManyBody<Collection>().strength((d) => r * d.items.length)
      )
      // Tick function of simulation
      .on("tick", ticked);

    // grab svg by ref (this element) and join with node
    const svgElement = d3
      .select(ref.current)
      .attr("width", width)
      .attr("height", height)
      .attr("viewBox", [-width / 2, -height / 2, width, height])
      .attr("style", "max-width: 100%; height: auto;")
      .attr("style", "border: solid");

    // Append circles to "this" scg element
    const node = svgElement
      .selectAll("circle")
      .data(nodes)
      .join(
        (enter) =>
          enter.append("circle").call((enter) =>
            enter
              .transition()
              .duration(500)
              .attr("r", (d) => r * d.items.length)
          ),
        (update) => update,
        (exit) => exit.transition().duration(500).attr("r", 0).remove()
      )
      .attr("fill", (d) => color(d.name));

    // Change node on tick
    function ticked() {
      node.attr("cx", (d) => d.x!).attr("cy", (d) => d.y!);
    }
  }, [nodes]);

  return (
    <>
      <div>
        <button onClick={Remove}>REMOVE</button>
        <br></br>
        <button onClick={Add}>ADD</button>
      </div>
      <svg viewBox="0 0 100 50" ref={ref} />
    </>
  );
};

export default Bubble;
