import java.util.Arrays;
import java.util.Locale;

public class LogLevels {

    public static void main(String[] args) {
        System.out.println(reformat("[ERROR]: Invalid operation"));
    }
    public static String message(String logLine) {
        String[] parts = {""};
        if (logLine.startsWith("[INFO]:")) {
            parts = logLine.split("\\[INFO\\]\\:");
        } else if (logLine.startsWith("[ERROR]:")) {
            parts = logLine.split("\\[ERROR\\]\\:");
        } else if (logLine.startsWith("[WARNING]:")) {
            parts = logLine.split("\\[WARNING\\]\\:");
        }
        return parts[1].trim();
    }

    public static String logLevel(String logLine) {
        String[] parts = {""};
        parts = logLine.split("\\]:");
        String logLevel = parts[0];
        logLevel = logLevel.substring(logLevel.indexOf("[")+1);
        return logLevel.trim().toLowerCase(Locale.ROOT);
    }

    public static String reformat(String logLine) {
        String[] parts = {""};
        parts = logLine.split("\\]:");
        String logLevel = parts[0];
        logLevel = logLevel.substring(logLevel.indexOf("[")+1).toLowerCase(Locale.ROOT);
        String message = parts[1].trim();
        return message+" ("+logLevel+")";
    }
}
