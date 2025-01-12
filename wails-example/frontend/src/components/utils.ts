export const isJson = (str: string) => {
  try {
    JSON.parse(str);
  } catch (e) {
    return false;
  }
  return true;
};

export const beautifyJson = (str: string, indent: number = 2) => {
  const obj = JSON.parse(str);
  return JSON.stringify(obj, null, indent);
};
