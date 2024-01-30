import { useRef, useEffect, useState } from "react";

import * as d3 from "d3";
import { SimulationNodeDatum } from "d3";

// Interface Node has...
interface Node {
  radius: () => number;
}

// Extends the Node type used in simulations
export type CollectionType = {
  id: number;
  name: string;
  items: object[];
} & SimulationNodeDatum;

// Bubble defines a function creating a bubble
// Change accepted children type to interface such that Bubble accepts
// generalised node children.
const Bubble = ({ children }: { children: Node[] }) => {
  const ref = useRef(null);

  const data = children.map((c) => {
    const col: CollectionType = { id: c.id, name: c.name, items: c.items };
    return col;
  });

  const [nodes, SetCol] = useState<CollectionType[]>(data);

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
          .forceCollide<CollectionType>((d) => r * d.items.length + 2)
          .iterations(12)
      )
      // Inter-node gravity
      .force(
        "charge",
        d3.forceManyBody<CollectionType>().strength((d) => r * d.items.length)
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
