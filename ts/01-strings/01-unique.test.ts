// import { expect, it } from "@jest/globals";
import { isUnique, isUniqueSimple, isUniqueSort } from "./01-unique";

it("isUnique", () => {
    const functions = [isUnique, isUniqueSimple, isUniqueSort];

    functions.forEach((func) => {
        expect(func("abc")).toBeTruthy();
        expect(func("abcd")).toBeTruthy();
        expect(func("abcd ")).toBeTruthy();
        expect(func("abca")).toBeFalsy();
        expect(func("abcc")).toBeFalsy();
        expect(func("a bc ")).toBeFalsy();
    });
});
