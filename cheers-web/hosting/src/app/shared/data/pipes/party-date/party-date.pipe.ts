import { Pipe, PipeTransform } from '@angular/core';

@Pipe({
      name: 'partyDate'
})
export class PartyDatePipe implements PipeTransform {
      transform(epoch: number) {
          const date = new Date(epoch);
          const today = new Date();
          const tomorrow = new Date(today.getTime() + 24 * 60 * 60 * 1000);

          const now = new Date();
          const isToday = date.toDateString() === now.toDateString();
          const isTomorrow = date.toDateString() === tomorrow.toDateString();

          const amPm = date.getHours() >= 12 ? 'PM' : 'AM';
          const time = date.getHours()+":"+String(date.getMinutes()).padStart(2, '0') + " " + amPm;

          if (isToday) {
              return `Today - ${time}`
          } else if (isTomorrow) {
              return `Tomorrow - ${time}`
          } else {
              const dayNames = ["Sun", "Mon", "Tue", "Wed", "Thu", "Fri", "Sat"];
              const months = ["Jan", "Feb", "Mar", "Apr", "May", "June", "July", "Aug", "Sept", "Oct", "Nov", "Dec"];

              return `${dayNames[date.getDay()]} ${date.getDate()} ${months[date.getMonth()]} - ${time}`;
          }
      }
}
