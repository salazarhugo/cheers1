import {Pipe, PipeTransform} from '@angular/core';

@Pipe({
    name: 'relativeTime'
})
export class RelativeTimePipe implements PipeTransform {
    transform(unix: number) {
        const date = new Date(unix * 1000)
        let seconds: number = Math.floor(((new Date()).getTime() - date.getTime()) / 1000);
        let interval: number = Math.floor(seconds / 31536000);

        if (interval > 1) {
            return interval + ' years ago';
        }
        interval = Math.floor(seconds / 2592000);
        if (interval > 1) {
            return interval + ' months ago';
        }
        interval = Math.floor(seconds / 86400);
        if (interval > 1) {
            return interval + ' days ago';
        }
        interval = Math.floor(seconds / 3600);
        if (interval > 1) {
            return interval + ' hours ago';
        }
        interval = Math.floor(seconds / 60);
        if (interval > 1) {
            return interval + ' minutes ago';
        }
        interval = Math.floor(seconds);
        if (interval > 10) {
            return interval + ' seconds ago';
        }

        return 'now';
    }
}
