import {beforeEach, describe, it, expect, jest} from '@jest/globals';
import { EventService, Event } from './event';
import {ApiService} from "./api";

describe('Test : ComicService + getComic', () => {
    let eventService: EventService;

    beforeEach(() => {
        eventService = new EventService();
    });

    it('should return a Promise of Comic when calling getComic', () => {
        const mockEvent: Event = {Id: 1};
        jest.spyOn(ApiService, 'getData').mockResolvedValue(mockEvent);

        // Call the getComic method
        const result = eventService.getComic(1);

        // Expect the result to be the mock comic
        //expect(result).toEqual(mockEvent);
        expect(result).toBeInstanceOf(Promise);
    });
});

describe('Test : ComicService + getAllComic', () => {
    let comicService: EventService;

    beforeEach(() => {
        comicService = new EventService();
    });

    it('should return a Promise of Comic when calling getAllComic', () => {
        const mockEvent: Event[] = [{Id: 1}];
        jest.spyOn(ApiService, 'getData').mockResolvedValue(mockEvent);

        const result = comicService.getAllComic();
        expect(result).toBeInstanceOf(Promise);
    });
});


describe('Test : ComicService + postComic', () => {
    let eventService: EventService;
    const event: Event = {Id: 1}

    beforeEach(() => {
        eventService = new EventService();
    });

    it('should return a Promise of Comic when calling postComic', () => {
        jest.spyOn(ApiService, 'postData').mockResolvedValue(1);

        const result = eventService.postComic(event);
        expect(result).toBeInstanceOf(Promise);
    });
});

describe('Test : ComicService + deleteComic', () => {
    let eventService: EventService;

    beforeEach(() => {
        eventService = new EventService();
    });

    it('should return a Promise of Comic when calling deleteComic', () => {
        jest.spyOn(ApiService, 'deleteData').mockResolvedValue(0);

        const result = eventService.deleteComic(1);
        expect(result).toBeInstanceOf(Promise);
    });
});

describe('Test : ComicService + putComic', () => {
    let eventService: EventService;
    const event: Event = {Id: 1}

    beforeEach(() => {
        eventService = new EventService();
    });

    it('should return a Promise of Comic when calling putComic', () => {
        jest.spyOn(ApiService, 'putData').mockResolvedValue(0);

        const result = eventService.putComic(event, 1);
        expect(result).toBeInstanceOf(Promise);
    });
});