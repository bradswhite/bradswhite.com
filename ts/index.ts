window.onload = () => {

  /**
    * Mode switcher
    * Gets mode status (light or dark) and displays the respective icon and switches mode when clicked
    */
  const switcher = document.getElementById("modeSwitcher");

  // false for light, true for dark
  var mode = localStorage.theme === 'dark' || (!('theme' in localStorage) && window.matchMedia('(prefers-color-scheme: dark)').matches);
  
  // Icons: (from Lucide icons)
  const sunIcon = '<svg xmlns="http://www.w3.org/2000/svg" width="32" height="32" viewBox="0 0 24 24"><g fill="none" stroke="#888888" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="4"/><path d="M12 2v2m0 16v2M4.93 4.93l1.41 1.41m11.32 11.32l1.41 1.41M2 12h2m16 0h2M6.34 17.66l-1.41 1.41M19.07 4.93l-1.41 1.41"/></g></svg>';
  const moonIcon = '<svg xmlns="http://www.w3.org/2000/svg" width="32" height="32" viewBox="0 0 24 24"><path fill="none" stroke="#888888" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 3a6 6 0 0 0 9 9a9 9 0 1 1-9-9"/></svg>';

  switcher.innerHTML = mode ? sunIcon : moonIcon;
  document.body.className = mode ? 'dark' : 'light';

  switcher.onclick = () => {
    if (mode) {
      switcher.innerHTML = moonIcon;
      mode = false;
      document.body.className = 'light';
      localStorage.theme = 'light';
    } else {
      switcher.innerHTML = sunIcon;
      mode = true;
      document.body.className = 'dark';
      localStorage.theme = 'dark';
    }
  };


  /**/
};
