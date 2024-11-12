components {
  id: "end"
  component: "/game/end.script"
}
embedded_components {
  id: "sprite"
  type: "sprite"
  data: "default_animation: \"cutting_board_64\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/assets/objects.atlas\"\n"
  "}\n"
  ""
}
embedded_components {
  id: "collisionobject"
  type: "collisionobject"
  data: "type: COLLISION_OBJECT_TYPE_TRIGGER\n"
  "mass: 0.0\n"
  "friction: 0.1\n"
  "restitution: 0.5\n"
  "group: \"end\"\n"
  "mask: \"player\"\n"
  "embedded_collision_shape {\n"
  "  shapes {\n"
  "    shape_type: TYPE_BOX\n"
  "    position {\n"
  "    }\n"
  "    rotation {\n"
  "    }\n"
  "    index: 0\n"
  "    count: 3\n"
  "  }\n"
  "  data: 68.69919\n"
  "  data: 21.290323\n"
  "  data: 10.0\n"
  "}\n"
  ""
}
embedded_components {
  id: "win_chef"
  type: "sprite"
  data: "default_animation: \"chef_knife_up\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/assets/objects.atlas\"\n"
  "}\n"
  ""
  position {
    x: 4.0
    y: 104.0
  }
  scale {
    x: 0.2
    y: 0.2
  }
}
