const preloaderElement = document.getElementsByTagName( 'preloader' )[0];

let preload = {
    init: () => {
        const bars = [];
        let i = 0;

        while( i < 3 ) {
            i++;
            const bar = document.createElement('bar');
            bars.push( bar );
            preloaderElement.appendChild( bar );
        }

        const barAnimation = function (index) {

            setTimeout(function () {

                setInterval(function () {

                    bars[index].setAttribute('class', (bars[index].getAttribute('class') === 'active') ? '' : 'active');
                }, 700);

            }, (index === 0) ? 50 : index * 150 + 50);
        };

        preloaderElement.setAttribute( 'class', 'animate' );

        setTimeout(function(){

            preloaderElement.setAttribute( 'class', 'start animate' );
        }, 300);
        setTimeout(function(){

            preloaderElement.setAttribute( 'class', 'start complete' );
        }, 1100);

        setTimeout(function(){
            for (let b = 0; b < bars.length; b++) {

                barAnimation( b );
            }
        }, 1100);
    },
    show: () => {
        preloaderElement.show = function() {
            this.style.display = 'block';
        };
    },
    hide: () => {
        preloaderElement.hide = function() {
            this.style.display = 'none';
        };
    }
}