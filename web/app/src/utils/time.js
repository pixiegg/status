/**
 * Generates a human-readable relative time string (e.g., "2 hours ago")
 * @param {string|Date} timestamp - The timestamp to convert
 * @returns {string} Relative time string
 */
export const generatePrettyTimeAgo = (timestamp) => {
  let differenceInMs = new Date().getTime() - new Date(timestamp).getTime();
  if (differenceInMs < 500) {
    return "agora";
  }
  if (differenceInMs > 3 * 86400000) {
    // If it was more than 3 days ago, we'll display the number of days ago
    let days = (differenceInMs / 86400000).toFixed(0);
    return days + " dia" + (days !== "1" ? "s" : "") + " atrás";
  }
  if (differenceInMs > 3600000) {
    // If it was more than 1h ago, display the number of hours ago
    let hours = (differenceInMs / 3600000).toFixed(0);
    return hours + " hora" + (hours !== "1" ? "s" : "") + " atrás";
  }
  if (differenceInMs > 60000) {
    let minutes = (differenceInMs / 60000).toFixed(0);
    return minutes + " minuto" + (minutes !== "1" ? "s" : "") + " atrás";
  }
  let seconds = (differenceInMs / 1000).toFixed(0);
  return seconds + " segundo" + (seconds !== "1" ? "s" : "") + " atrás";
};

/**
 * Generates a pretty time difference string between two timestamps
 * @param {string|Date} start - Start timestamp
 * @param {string|Date} end - End timestamp
 * @returns {string} Time difference string
 */
export const generatePrettyTimeDifference = (start, end) => {
  const ms = new Date(start) - new Date(end);
  const seconds = Math.floor(ms / 1000);
  const minutes = Math.floor(seconds / 60);
  const hours = Math.floor(minutes / 60);

  if (hours > 0) {
    const remainingMinutes = minutes % 60;
    const hoursText = hours + (hours === 1 ? " hora" : " horas");
    if (remainingMinutes > 0) {
      return (
        hoursText +
        " " +
        remainingMinutes +
        (remainingMinutes === 1 ? " minuto" : " minutos")
      );
    }
    return hoursText;
  } else if (minutes > 0) {
    const remainingSeconds = seconds % 60;
    const minutesText = minutes + (minutes === 1 ? " minuto" : " minutos");
    if (remainingSeconds > 0) {
      return (
        minutesText +
        " " +
        remainingSeconds +
        (remainingSeconds === 1 ? " segundo" : " segundos")
      );
    }
    return minutesText;
  } else {
    return seconds + (seconds === 1 ? " segundo" : " segundos");
  }
};

export const prettifyTimestamp = (timestamp) => {
  let date = new Date(timestamp);
  let YYYY = date.getFullYear();
  let MM = String(date.getMonth() + 1).padStart(2, "0");
  let DD = String(date.getDate()).padStart(2, "0");
  let hh = String(date.getHours()).padStart(2, "0");
  let mm = String(date.getMinutes()).padStart(2, "0");
  return `${DD}/${MM}/${YYYY} ${hh}:${mm}`;
};
