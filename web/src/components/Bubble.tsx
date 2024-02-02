import { useRef, useLayoutEffect, useState } from "react";

import * as d3 from "d3";
import { SimulationNodeDatum } from "d3";

/**
 * Node is a helper interface for rendering D3 Simulations with React & Typescript.
 *
 * This interface extends the D3.js type SimulationNodeDatum that is used by D3
 * simulations to dynamically update the positions of elements in the DOM. The
 * new required parameters allow for custom grouping of Nodes & defining how
 * the radii of Nodes are calculated.
 */
interface Node extends SimulationNodeDatum {
  group: string;
  radius: number;
}
/**
 * Bubble defines a component that renders a D3.js powered Bubble Plot given
 * children elements that satisfy the Node interface.
 *
 */
const Bubble = ({ children }: { children: Node[] }) => {
  // divRef is used to reference the plot's container and the changing width
  const divRef = useRef(null);
  // svgRef is used to reference the d3 svg itself and is necessary as both
  // React and D3 both manipulate the DOM.
  const svgRef = useRef(null);

  // Set up nodes of the Bubble Plot
  const nodes = children.map((c) => ({ ...c }));
  // const [nodes, SetNodes] = useState<Node[]>(children);

  // Establish width and height to later set by Ref.
  const [width, SetWidth] = useState(800);
  const [height, SetHeight] = useState(800);

  const r = 20;

  // const Remove = () => {
  //   const temp = nodes.map((d) => ({ ...d }));
  //   temp.pop();
  //   SetCol(temp);
  // };

  // const Add = () => {
  //   const temp = nodes.map((d) => ({ ...d }));
  //   const addThis = { group: 99, radius: 4 };
  //   SetCol(temp.concat(addThis));
  // };

  useLayoutEffect(() => {
    // Save current element width and height (TODO: Not yet Responsive)
    SetWidth(divRef.current ? divRef.current.offsetWidth : 800);
    SetHeight(divRef.current ? divRef.current.offsetHeight : 800);

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
        d3.forceCollide<Node>((d) => r * d.radius + 2).iterations(12)
      )
      // Inter-node gravity
      .force(
        "charge",
        d3.forceManyBody<Node>().strength((d) => r * d.radius)
      )
      .on("tick", ticked);

    // Establish SVG sizing and ViewBox
    const svgElement = d3
      .select(svgRef.current)
      .attr("width", width)
      .attr("height", height)
      .attr("viewBox", [-width / 2, -height / 2, width, height]);

    // Join Node Data to simulation as circles
    const node = svgElement
      .selectAll("circle")
      .data(nodes)
      .join(
        (enter) =>
          enter.append("circle").call((enter) =>
            enter
              .transition()
              .duration(500)
              .attr("r", (d) => r * d.radius)
          ),
        (update) => update,
        (exit) => exit.transition().duration(500).attr("r", 0).remove()
      )
      .attr("fill", (d) => color(d.group));

    // Change node on tick
    function ticked() {
      node.attr("cx", (d) => d.x!).attr("cy", (d) => d.y!);
    }
  }, [nodes, children, width, height]);

  return (
    <div className="h-full w-full" ref={divRef}>
      {/* <div>
        <button onClick={Remove}>REMOVE</button>
        <br></br>
        <button onClick={Add}>ADD</button>
      </div> */}
      <svg className="m-auto" ref={svgRef} />
    </div>
  );
};

export default Bubble;
