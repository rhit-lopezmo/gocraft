# Gocraft Roadmap  
*A high-level plan for building the voxel engine. No implementation details.*

---

## PHASE 1 — Project Bootstrap
- Set up project structure.
- Initialize a window and basic 3D camera.
- Create a minimal render loop.
- Add basic player movement in 3D space.

---

## PHASE 2 — Voxel Data System
- Define block types.
- Define chunk dimensions and data layout.
- Create a world container to hold chunks.
- Add ability to populate chunks with simple test data.

---

## PHASE 3 — Mesh Generation
- Implement a system to turn chunk data into renderable meshes.
- Add a dedicated asynchronous mesh generation pipeline.
- Process chunk mesh jobs in background workers.
- Update rendered models when mesh results are returned.

---

## PHASE 4 — World Loading Pipeline
- Determine visible chunks based on player position.
- Maintain a set of loaded chunks around the player.
- Add asynchronous world generation for new chunks.
- Queue generated chunks for meshing when ready.

---

## PHASE 5 — Player Interaction
- Add basic collision and grounded movement.
- Implement block selection using a raycast.
- Support placing and removing blocks.
- Trigger chunk remeshing when blocks are changed.

---

## PHASE 6 — World Expansion
- Add simple terrain generation (noise-based).
- Enable infinite or semi-infinite world streaming.
- Add chunk caching or unloading for out-of-range areas.
- Improve chunk scheduling and prioritization.

---

## PHASE 7 — Rendering Enhancements
- Improve chunk meshing quality.
- Add simple lighting or shading.
- Add frustum culling or distance-based culling.
- Optimize draw calls and memory usage.

---

## PHASE 8 — Polish & Performance
- Add configuration settings (render distance, FOV, sensitivity).
- Improve multithreading and reduce stalls.
- Profile CPU and GPU usage.
- Prepare builds for release and testing.

