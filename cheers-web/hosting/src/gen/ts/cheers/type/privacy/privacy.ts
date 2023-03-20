/* eslint-disable */

export const protobufPackage = "cheers.type";

export enum Privacy {
  FRIENDS = 0,
  PRIVATE = 1,
  PUBLIC = 2,
  GROUP = 3,
  UNRECOGNIZED = -1,
}

export function privacyFromJSON(object: any): Privacy {
  switch (object) {
    case 0:
    case "FRIENDS":
      return Privacy.FRIENDS;
    case 1:
    case "PRIVATE":
      return Privacy.PRIVATE;
    case 2:
    case "PUBLIC":
      return Privacy.PUBLIC;
    case 3:
    case "GROUP":
      return Privacy.GROUP;
    case -1:
    case "UNRECOGNIZED":
    default:
      return Privacy.UNRECOGNIZED;
  }
}

export function privacyToJSON(object: Privacy): string {
  switch (object) {
    case Privacy.FRIENDS:
      return "FRIENDS";
    case Privacy.PRIVATE:
      return "PRIVATE";
    case Privacy.PUBLIC:
      return "PUBLIC";
    case Privacy.GROUP:
      return "GROUP";
    case Privacy.UNRECOGNIZED:
    default:
      return "UNRECOGNIZED";
  }
}
