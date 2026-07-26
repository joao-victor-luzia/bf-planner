<script lang="ts">
	import Calendar from './Calendar.svelte';
	import Details from './Details.svelte';
	import type { CalendarTask, CalendarDay } from './types.ts';

	var dayNames = ['Sun', 'Mon', 'Tue', 'Wed', 'Thu', 'Fri', 'Sat'];
	let monthNames = [
		'January',
		'February',
		'March',
		'April',
		'May',
		'June',
		'July',
		'August',
		'September',
		'October',
		'November',
		'December'
	];

	let selectedDay: CalendarDay = $state({
		name: '',
		enabled: true,
		date: new Date()
	});
	let headers: string[] = dayNames;
	let now = new Date();
	let year = $state(now.getFullYear());
	let month = $state(now.getMonth());
	let eventText = $state('Click an item or date');

	var days: CalendarDay[] = $derived(initMonth(year, month));

	function randInt(max: number): number {
		return Math.floor(Math.random() * max) + 1;
	}

	var items: CalendarTask[] = $derived(initMonthItems(year, month));

	function initMonthItems(y: number, m: number) {
		let d1 = new Date(y, m, randInt(7) + 7);
		let itemsproc: CalendarTask[] = [];
		let items: CalendarTask[] = [
			{
				title: '11:00 Task Early in month',
				className: 'task--primary',
				date: new Date(y, m, randInt(6)),
				len: randInt(4) + 1
			},
			{ title: '7:30 Wk 2 tasks', className: 'task--warning', date: d1, len: randInt(4) + 2 },
			{
				title: 'Overlapping Stuff (isBottom:true)',
				date: d1,
				className: 'task--info',
				len: 4,
				isBottom: true
			},
			{
				title: '10:00 More Stuff to do',
				date: new Date(y, m, randInt(7) + 14),
				className: 'task--info',
				len: randInt(4) + 1,
				detailHeader: 'Difficult',
				detailContent: 'But not especially so'
			},
			{
				title: 'All day task',
				date: new Date(y, m, randInt(7) + 21),
				className: 'task--danger',
				len: 1,
				vlen: 2
			}
		];

		for (let i of items) {
			let rc = findRowCol(i.date);
			if (rc == null) {
				console.log('didn`t find date for ', i);
				console.log(i.date);
				console.log(days);
				i.startCol = i.startRow = 0;
			} else {
				i.startCol = rc.col;
				i.startRow = rc.row;
				if (i.startCol + i.len > 8 && i.startRow < 6) {
					let copyi = structuredClone(i);
					copyi.startCol = 1;
					copyi.startRow!++;
					copyi.len = i.startCol + i.len - 8;
					itemsproc.push(copyi);
				}
			}
		}
		return [...items, ...itemsproc];
	}

	function initMonth(year: number, month: number) {
		let days = [];
		let monthAbbrev = monthNames[month].slice(0, 3);
		let nextMonthAbbrev = monthNames[(month + 1) % 12].slice(0, 3);
		var firstDay = new Date(year, month, 1).getDay();
		var daysInThisMonth = new Date(year, month + 1, 0).getDate();
		var daysInLastMonth = new Date(year, month, 0).getDate();
		var prevMonth = month == 0 ? 11 : month - 1;

		for (let i = daysInLastMonth - firstDay; i < daysInLastMonth; i++) {
			let d = new Date(prevMonth == 11 ? year - 1 : year, prevMonth, i + 1);
			days.push({ name: '' + (i + 1), enabled: false, date: d });
		}
		for (let i = 0; i < daysInThisMonth; i++) {
			let d = new Date(year, month, i + 1);
			if (i == 0) days.push({ name: monthAbbrev + ' ' + (i + 1), enabled: true, date: d });
			else days.push({ name: '' + (i + 1), enabled: true, date: d });
		}
		for (let i = 0; days.length % 7; i++) {
			let d = new Date(month == 11 ? year + 1 : year, (month + 1) % 12, i + 1);
			if (i == 0) days.push({ name: nextMonthAbbrev + ' ' + (i + 1), enabled: false, date: d });
			else days.push({ name: '' + (i + 1), enabled: false, date: d });
		}
		return days;
	}

	function findRowCol(dt: Date) {
		for (let i = 0; i < days.length; i++) {
			let d = days[i].date;
			if (
				d.getFullYear() === dt.getFullYear() &&
				d.getMonth() === dt.getMonth() &&
				d.getDate() === dt.getDate()
			) {
				return { row: Math.floor(i / 7) + 2, col: (i % 7) + 1 };
			}
		}
		return null;
	}

	function itemClick(e: CalendarTask) {
		eventText = 'itemClick ' + JSON.stringify(e) + ' localtime=' + e.date.toString();
	}
	function dayClick(e: CalendarDay) {
		eventText = 'onDayClick ' + JSON.stringify(e) + ' localtime=' + e.date.toString();
		selectedDay = e;
	}
	function headerClick(e: string) {
		eventText = 'onHheaderClick ' + JSON.stringify(e);
	}
	function next() {
		month++;
		if (month == 12) {
			year++;
			month = 0;
		}
	}
	function prev() {
		if (month == 0) {
			month = 11;
			year--;
		} else {
			month--;
		}
	}
</script>

<main>
	<div class="calendar-container">
		<div class="calendar-header">
			<h1>
				<button onclick={() => year--}>&Lt;</button>
				<button onclick={() => prev()}>&lt;</button>
				{monthNames[month]}
				{year}
				<button onclick={() => next()}>&gt;</button>
				<button onclick={() => year++}>&Gt;</button>
			</h1>
			{eventText}
		</div>

		<Calendar
			{headers}
			{days}
			{items}
			onClickDay={(e: CalendarDay) => dayClick(e)}
			onClickItem={(e: CalendarTask) => itemClick(e)}
			onClickHeader={(e: string) => headerClick(e)}
		/>
	</div>
	<div class="event-container">
		<Details {eventText} {selectedDay} />
	</div>
</main>

<style>
	main {
		display: flex;
		flex-direction: row;
		justify-content: space-between;
		width: 100%;
	}
	.calendar-container {
		width: 70%;
		overflow: hidden;
		box-shadow: 0 2px 20px rgba(0, 0, 0, 0.1);
		border-radius: 10px;
		background: #fff;
	}
	.event-container {
		width: 29%;
		overflow: hidden;
		box-shadow: 0 2px 20px rgba(0, 0, 0, 0.1);
		border-radius: 10px;
		background: #fff;
	}
	.calendar-header {
		text-align: center;
		padding: 20px 0;
		background: #eef;
		border-bottom: 1px solid rgba(166, 168, 179, 0.12);
	}
	.calendar-header h1 {
		margin: 0;
		font-size: 18px;
	}
	.calendar-header button {
		background: #eef;
		border: 1px;
		padding: 6;
		color: rgba(81, 86, 93, 0.7);
		cursor: pointer;
		outline: 0;
	}
</style>
