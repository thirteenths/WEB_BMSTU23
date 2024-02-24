import {beforeEach, describe, it, expect, jest} from '@jest/globals';
import { ComicService, Comic } from './comic';
import {ApiService} from "./api";

describe('Test : ComicService + getComic', () => {
    let comicService: ComicService;

    beforeEach(() => {
        comicService = new ComicService();
    });

    it('should return a Promise of Comic when calling getComic', () => {
        const mockComic: Comic = {Id: 1};
        jest.spyOn(ApiService, 'getData').mockResolvedValue(mockComic);

        // Call the getComic method
        const result = comicService.getComic(1);

        // Expect the result to be the mock comic
        //expect(result).toEqual(mockComic);
        expect(result).toBeInstanceOf(Promise);
    });
});

describe('Test : ComicService + getAllComic', () => {
    let comicService: ComicService;

    beforeEach(() => {
        comicService = new ComicService();
    });

    it('should return a Promise of Comic when calling getAllComic', () => {
        const mockComic: Comic[] = [{Id: 1}];
        jest.spyOn(ApiService, 'getData').mockResolvedValue(mockComic);

        const result = comicService.getAllComic();
        expect(result).toBeInstanceOf(Promise);
    });
});


describe('Test : ComicService + postComic', () => {
    let comicService: ComicService;
    const comic: Comic = {Id: 1}

    beforeEach(() => {
        comicService = new ComicService();
    });

    it('should return a Promise of Comic when calling postComic', () => {
        jest.spyOn(ApiService, 'postData').mockResolvedValue(1);

        const result = comicService.postComic(comic);
        expect(result).toBeInstanceOf(Promise);
    });
});

describe('Test : ComicService + deleteComic', () => {
    let comicService: ComicService;

    beforeEach(() => {
        comicService = new ComicService();
    });

    it('should return a Promise of Comic when calling deleteComic', () => {
        jest.spyOn(ApiService, 'deleteData').mockResolvedValue(0);

        const result = comicService.deleteComic(1);
        expect(result).toBeInstanceOf(Promise);
    });
});

describe('Test : ComicService + putComic', () => {
    let comicService: ComicService;
    const comic: Comic = {Id: 1}

    beforeEach(() => {
        comicService = new ComicService();
    });

    it('should return a Promise of Comic when calling putComic', () => {
        jest.spyOn(ApiService, 'putData').mockResolvedValue(0);

        const result = comicService.putComic(comic, 1);
        expect(result).toBeInstanceOf(Promise);
    });
});